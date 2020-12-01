package limiter

import (
	"net"
	"net/http"

	"github.com/collabyt/Backend/cache"
	"github.com/collabyt/Backend/handler"
	"github.com/collabyt/Backend/logger"
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
			logger.Error.Println("Error Splitting IP address and Port of Client")
			handler.WriteErrorReply(w, http.StatusInternalServerError)
			return
		}
		ctx := r.Context()
		limiter := redis_rate.NewLimiter(cache.Cache)
		rRet, err := limiter.Allow(ctx, ip, redis_rate.PerMinute(cache.CacheTTL))
		if err != nil {
			logger.Error.Println("Error processing Limiter permission")
			handler.WriteErrorReply(w, http.StatusInternalServerError)
			return

		}
		if rRet.Allowed == 0 {
			logger.Error.Printf("Connection attempt by %s denied!\n", r.RemoteAddr)
			handler.WriteErrorReply(w, http.StatusTooManyRequests)
			return
		}
		logger.Info.Printf("New HTTP request from %s; Total allowed Requests %d of %d\n", r.RemoteAddr, rRet.Allowed, rRet.Remaining)
		next.ServeHTTP(w, r)
	})
}
