package limiter

import (
	"fmt"
	"net"
	"net/http"

	"github.com/collabyt/Backend/cache"
	"github.com/collabyt/Backend/handler"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

// Limit cap the amount of requests possible for a single IP in a given period
// of time
func Limit(rClient *redis.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			handler.WriteErrorReply(w, http.StatusInternalServerError)
			return
		}
		ctx := r.Context()
		limiter := redis_rate.NewLimiter(cache.Cache)
		rRet, err := limiter.Allow(ctx, ip, redis_rate.PerMinute(timeToDumpVisitor))
		if err != nil {
			handler.WriteErrorReply(w, http.StatusInternalServerError)
			return

		}
		if rRet.Allowed == 0 {
			handler.WriteErrorReply(w, http.StatusTooManyRequests)
			return
		}
		fmt.Println("allowed", rRet.Allowed, "remaining", rRet.Remaining)
		next.ServeHTTP(w, r)
	})
}

/*
// Visitor is a instance of a user accessing the API
type Visitor struct {
	limiter *rate.Limiter
	expiration
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
*/
