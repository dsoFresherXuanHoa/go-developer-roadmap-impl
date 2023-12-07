package gogin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouteConfig(router *gin.Engine, db *gorm.DB) {
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/sign-up", SignUp(db))
	}
}
