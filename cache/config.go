package cache

import "os"

var (
	cacheHost     = os.Getenv("CACHE_HOST")
	cachePort     = os.Getenv("CACHE_PORT")
	cachePassword = os.Getenv("CACHE_PASSWORD")
	cacheDB       = 0
	CacheTTL      = 0
)
