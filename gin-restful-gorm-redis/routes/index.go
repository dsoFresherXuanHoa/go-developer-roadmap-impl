package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB, redisClient *redis.Client) {
	router := gin.Default()

	AccountRouteConfig(db, router, redisClient)

	router.Run()
}
