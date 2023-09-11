package routes

import (
	"gin-restful-gorm/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContactRouteConfig(db *gorm.DB, router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		contacts := v1.Group("/contacts")
		{
			contacts.GET("/", controllers.FindAllContact(db))
			contacts.GET("/:id", controllers.FindContactByID(db))
			contacts.POST("/", controllers.SaveContact(db))
			contacts.PUT("/", controllers.UpdateContact(db))
			contacts.PATCH("/", controllers.UpdateContact(db))
			contacts.DELETE("/:id", controllers.DeleteContact(db))
		}
	}
}
