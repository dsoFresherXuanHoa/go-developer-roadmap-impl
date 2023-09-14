package configs

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func RedisConfig() *redis.Client {
	if err := godotenv.Load(); err != nil {
		panic("Can't load .env variable!")
	}

	var REDIS_NETWORK = os.Getenv("REDIS_NETWORK")

	return redis.NewClient(&redis.Options{
		Addr:     REDIS_NETWORK + ":6379",
		Password: "",
		DB:       0,
	})
}
