package routes

import (
	"gin-restful-gorm-redis/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func AccountRouteConfig(db *gorm.DB, router *gin.Engine, redisClient *redis.Client) {
	accounts := router.Group("/accounts")
	{
		accounts.GET("/", controllers.FindAllAccount(db))
		accounts.GET("/:username", controllers.FindAccountByUsername(db, redisClient))
		accounts.POST("/", controllers.SaveAccount(db, redisClient))
		accounts.PUT("/", controllers.UpdateAccount(db, redisClient))
		accounts.PATCH("/", controllers.UpdateAccount(db, redisClient))
		accounts.DELETE("/", controllers.DeleteAccount(db, redisClient))
	}
}
