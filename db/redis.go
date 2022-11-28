package db

import (
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func CnnectRedis() *redis.Client {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}
