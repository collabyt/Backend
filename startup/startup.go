package startup

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/collabyt/Backend/cache"
	"github.com/collabyt/Backend/limiter"
	"github.com/gorilla/mux"
)

func SetupServer(r *mux.Router) (http.Server, error) {
	address := os.Getenv("APP_PORT")
	handler := limiter.Limit(cache.Cache, r)
	idleTimeout, err := strconv.Atoi(os.Getenv("APP_IDLE_TIMEOUT"))
	if err != nil {
		return http.Server{}, fmt.Errorf("server stup faile: impossible to get %s from system's parameters", "APP_IDLE_TIMEOUT")
	}
	readTimeout, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	if err != nil {
		return http.Server{}, fmt.Errorf("server stup faile: impossible to get %s from system's parameters", "APP_READ_TIMEOUT")
	}
	writeTimeout, err := strconv.Atoi(os.Getenv("APP_WRITE_TIMEOUT"))
	if err != nil {
		return http.Server{}, fmt.Errorf("server stup faile: impossible to get %s from system's parameters", "APP_WRITE_TIMEOUT")
	}
	return http.Server{
		Addr:         address,
		Handler:      handler,
		IdleTimeout:  time.Duration(idleTimeout),
		ReadTimeout:  time.Duration(readTimeout),
		WriteTimeout: time.Duration(writeTimeout),
	}, err
}
