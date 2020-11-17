package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"

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
	rClient := redis.NewClient(rOptions)
	ctx := context.Background()
	_, err = rClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	_ = rClient.FlushDB(ctx).Err()
	Cache = rClient
}

func loadTTL() (int, error) {
	ttl, err := strconv.Atoi(os.Getenv("CACHE_TTL"))
	return ttl, err
}
