package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RouterConfig(client *mongo.Client) {
	router := gin.Default()

	ContactRouterConfig(router, client)

	router.Run()
}
