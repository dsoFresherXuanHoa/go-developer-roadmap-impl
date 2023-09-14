package controllers

import (
	"gin-restful-gorm-redis/models"
	"gin-restful-gorm-redis/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func FindAllAccount(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if contacts, err := services.FindAllAccount(db); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if len(contacts) == 0 {
			ctx.JSON(http.StatusNoContent, gin.H{
				"message": "No record found!",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": contacts,
			})
		}
	}
}

func FindAccountByUsername(db *gorm.DB, redisClient *redis.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		if account, err := services.FindAccountByUsername(username, db, redisClient); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "No record found!",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": account,
			})
		}
	}
}

func SaveAccount(db *gorm.DB, redisClient *redis.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var account models.Account
		if err := ctx.ShouldBind(&account); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "JSON Parsing Error: " + err.Error(),
			})
		} else if err := services.SaveAccount(db, account, redisClient); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Can't save record to database: " + err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Record has been saved!",
			})
		}
	}
}

func UpdateAccount(db *gorm.DB, redisClient *redis.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var account models.AccountUpdatable
		if err := ctx.ShouldBind(&account); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "JSON Parsing Error: " + err.Error(),
			})
		} else {
			username := ctx.Query("username")
			if err := services.UpdateAccount(db, username, account, redisClient); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Can't save record to database: " + err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Record has been update!",
				})
			}
		}
	}
}

func DeleteAccount(db *gorm.DB, redisClient *redis.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		username := ctx.Query("username")
		if err := services.DeleteAccount(db, username, redisClient); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Can't deleted record from database: " + err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Record has been deleted!",
			})
		}
	}
}
