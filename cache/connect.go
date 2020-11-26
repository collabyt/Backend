package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/collabyt/Backend/logger"
	"github.com/go-redis/redis/v8"
)

var (
	// Cache is the REDIS document database
	Cache *redis.Client
)

// Connect will initiate a connection with the 0 DB on the REDIS server
func Connect() {
	ttl, err := loadTTL()
	if err != nil {
		panic(err)
	}
	CacheTTL = ttl
	rOptions := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cacheHost, cachePort),
		Password: cachePassword,
		DB:       0,
	}
	logger.Info.Println("Trying to connect to REDIS cache server...")
	rClient := redis.NewClient(rOptions)
	ctx := context.Background()
	_, err = rClient.Ping(ctx).Result()
	if err != nil {
		logger.Error.Println("Connection to REDIS cache server failed!")
		panic(err)
	}

	_ = rClient.FlushDB(ctx).Err()
	Cache = rClient
	logger.Info.Println("Successfully connected to REDIS cache server")
}

func loadTTL() (int, error) {
	ttl, err := strconv.Atoi(os.Getenv("CACHE_TTL"))
	return ttl, err
}
