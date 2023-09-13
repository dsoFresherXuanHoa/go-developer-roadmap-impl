package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB) {
	router := gin.Default()

	ContactRouteConfig(db, router)

	router.Run(":3001")
}
