package limiter

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/collabyt/Backend/handler"
)

// Visitor is a instance of a user accessing the API
type Visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Catalog represent a list of visitor and a Mutex
type Catalog struct {
	visitors map[string]*Visitor
	mu       sync.RWMutex
	rate     rate.Limit
	maxBurst int
}

// NewCatalog initialize the visitors map inside the type Catalog
func NewCatalog(size, r, maxBurst int) *Catalog {
	var cat Catalog
	cat.visitors = make(map[string]*Visitor, size)
	cat.rate = rate.Limit(r)
	cat.maxBurst = maxBurst
	return &cat
}

func getVisitor(cat *Catalog, ip string) *rate.Limiter {
	cat.mu.RLock() //read
	v, exist := cat.visitors[ip]
	cat.mu.RUnlock()
	if !exist {
		limiter := rate.NewLimiter(cat.rate, cat.maxBurst)
		cat.mu.Lock() //read
		cat.visitors[ip] = &Visitor{limiter, time.Now()}
		cat.mu.Unlock()
		return limiter
	}
	cat.mu.Lock() //write
	v.lastSeen = time.Now()
	cat.mu.Unlock()
	return v.limiter
}

// Limit cap the amount of requests possible for a single IP in a given period
// of time
func Limit(cat *Catalog, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			handler.WriteErrorReply(w, http.StatusInternalServerError)
			return
		}
		limiter := getVisitor(cat, ip)
		if !limiter.Allow() {
			handler.WriteErrorReply(w, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CleanupVisitors remove unused IP's from the visitors
func CleanupVisitors(cat *Catalog, limit time.Duration) {
	for {
		time.Sleep(time.Minute)
		cat.mu.Lock()
		for ip, v := range cat.visitors {
			if time.Since(v.lastSeen) > limit {
				delete(cat.visitors, ip)
			}
		}
		cat.mu.Unlock()
	}
}
