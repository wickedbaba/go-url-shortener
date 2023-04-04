package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var pl = fmt.Println
var spf = fmt.Sprintf
var pf = fmt.Printf

func errChecker(err error) {
	if err != nil {
		panic(spf("Error init Redis: %v", err))
	}
}

// Defining the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// set only for this project
const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	errChecker(err)

	pf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

/* We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {

	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/

func RetrieveInitialUrl(shortUrl string) string {

	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
