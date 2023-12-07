package routes

import (
	"go-webcrawl-api/src/api/controllers"

	"github.com/gin-gonic/gin"
)

func NewsRouteConfig(router *gin.Engine) {
	goCrawl := router.Group("/")
	{
		goCrawl.POST("/crawl", controllers.SyncCrawlingNews())
		goCrawl.POST("/go-crawl", controllers.AsyncCrawlingNews())
	}
}
