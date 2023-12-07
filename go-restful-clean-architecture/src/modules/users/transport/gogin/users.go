package gogin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersRouteConfig(router *gin.Engine, db *gorm.DB) {
	users := router.Group("/api/v1/users")
	{
		users.POST("/", Create(db))
	}
}
