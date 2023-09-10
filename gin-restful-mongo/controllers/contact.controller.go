package controllers

import (
	"gin-restful-mongo/models"
	"gin-restful-mongo/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllContact(client *mongo.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if contacts, err := services.FindAllContact(ctx, client); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "No record found: " + err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": contacts,
			})
		}
	}
}

func FindContactByName(client *mongo.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var name = ctx.Param("name")
		if contact, err := services.FindContactByName(ctx, client, name); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "No record find!",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": contact,
			})
		}
	}
}

func SaveContact(client *mongo.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var contact models.Contact
		if err := ctx.ShouldBind(&contact); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			if result, err := services.SaveContact(ctx, client, contact); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Can't save record to database!",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Record has been saved: ",
					"data":    result,
				})
			}
		}
	}
}

func UpdateContact(client *mongo.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var contact models.Contact
		if err := ctx.ShouldBind(&contact); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			name := ctx.Query("name")
			if result, err := services.UpdateContact(ctx, client, contact, name); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Can't update record to database!",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Record has been saved!",
					"data":    result,
				})
			}
		}
	}
}

func DeleteContact(client *mongo.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		if result, err := services.DeleteContact(ctx, client, name); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Can't deleted record to database!",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Record has been deleted!",
				"data":    result,
			})
		}
	}
}
