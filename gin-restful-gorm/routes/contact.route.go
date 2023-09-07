package routes

import (
	"gin-restful-gorm/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContactRouteConfig(db *gorm.DB, router *gin.Engine) {
	v1 := router.Group("/v1/contacts/")
	{
		v1.GET("/", controllers.FindAllContact(db))
		v1.GET("/:id", controllers.FindContactByID(db))
		// Wrong response
		v1.POST("/", controllers.SaveContact(db))
		v1.PUT("/", controllers.UpdateContact(db))

		v1.DELETE("/:id", controllers.DeleteContact(db))
	}
}
