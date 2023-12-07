package routes

import "github.com/gin-gonic/gin"

func RouteConfig() {
	router := gin.Default()

	NewsRouteConfig(router)

	router.Run(":3000")
}
