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
	"github.com/rs/cors"
)

func SetupServer(r *mux.Router) (http.Server, error) {
	address := fmt.Sprintf("%s:%s", os.Getenv("APP_ADDRESS"), os.Getenv("APP_PORT"))
	handler := limiter.Limit(cache.Cache, r)
	corsAddress := os.Getenv("CORS_ADDRESS")
	corsPort := os.Getenv("CORS_PORT")
	corsHeader := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080", fmt.Sprintf("http://%s:%s", corsAddress, corsPort)},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})
	handler = corsHeader.Handler(handler)
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
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}, err
}
