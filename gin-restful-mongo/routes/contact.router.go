package routes

import (
	"gin-restful-mongo/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ContactRouterConfig(router *gin.Engine, client *mongo.Client) {
	v1 := router.Group("/v1")
	{
		contact := v1.Group("/contacts")
		{

			contact.GET("/", controllers.FindAllContact(client))
			contact.GET("/:name", controllers.FindContactByName(client))
			contact.POST("/", controllers.SaveContact(client))
			contact.PUT("/", controllers.UpdateContact(client))
			contact.DELETE("/:name", controllers.DeleteContact(client))
		}
	}
}
