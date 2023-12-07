package controllers

import (
	"context"
	"fmt"
	"go-webcrawl-api/src/api/constants"
	"go-webcrawl-api/src/api/services"
	"go-webcrawl-api/src/api/utils"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func SyncCrawlingNews() func(*gin.Context) {
	return func(ctx *gin.Context) {
		if urlSet, err := utils.UrlsFromSiteMap(constants.SITEMAP_FILEPATH); err != nil {
			fmt.Println(urlSet)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"data":    nil,
				"message": constants.ErrLoadSiteMap,
				"error":   err,
			})
		} else if err := services.SyncSaveNews(context.Background(), *urlSet); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":    nil,
				"message": constants.ErrSaveNews,
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data":    len(urlSet.Urls),
				"message": constants.Success,
				"error":   nil,
			})
		}
	}
}

func AsyncCrawlingNews() func(*gin.Context) {
	numberOfWorkers := 10
	wg := new(sync.WaitGroup)

	return func(ctx *gin.Context) {
		if urlSet, err := utils.UrlsFromSiteMap(constants.SITEMAP_FILEPATH); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"data":    nil,
				"message": constants.ErrLoadSiteMap,
				"error":   err,
			})
		} else {
			services.AsyncSaveNews(*urlSet, numberOfWorkers, wg)
			wg.Wait()
			ctx.JSON(http.StatusOK, gin.H{
				"data":    len(urlSet.Urls),
				"message": constants.Success,
				"error":   nil,
			})
		}
	}
}
