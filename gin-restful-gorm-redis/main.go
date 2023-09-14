package main

import (
	"gin-restful-gorm-redis/configs"
	"gin-restful-gorm-redis/routes"
	"log"
)

func main() {
	if db, err := configs.GormConfig(); err != nil {
		panic("Can't connect to database via GORM: " + err.Error())
	} else {
		redisClient := configs.RedisConfig()
		if _, err := redisClient.Ping().Result(); err != nil {
			panic("Can't connect to Redis Server: " + err.Error())
		} else {
			log.Println("Redis Server connected!")
			routes.RouteConfig(db, redisClient)
		}
	}
}
