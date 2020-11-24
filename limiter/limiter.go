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
		rRet, err := limiter.Allow(ctx, ip, redis_rate.PerMinute(cache.CacheTTL))
		if err != nil {
			handler.WriteErrorReply(w, http.StatusInternalServerError)
			return

		}
		if rRet.Allowed == 0 {
			handler.WriteErrorReply(w, http.StatusTooManyRequests)
			return
		}
		fmt.Println("allowed", rRet.Allowed, "remaining", rRet.Remaining) //TODO: Implement log message instead of println
		next.ServeHTTP(w, r)
	})
}
