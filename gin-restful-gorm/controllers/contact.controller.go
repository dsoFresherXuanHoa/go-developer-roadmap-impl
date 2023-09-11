package controllers

import (
	"gin-restful-gorm/models"
	"gin-restful-gorm/services"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAllContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if contacts, err := services.FindAllContact(db); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if (reflect.DeepEqual(contacts, models.Contacts{})) {
			ctx.JSON(http.StatusNoContent, gin.H{
				"message": "No record find!",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": contacts,
			})
		}
	}
}

func FindContactByID(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if id, err := strconv.Atoi(ctx.Param("id")); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if contact, err := services.FindContactById(id, db); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if contact == (models.Contact{}) {
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

func SaveContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var contact models.Contact
		if err := ctx.ShouldBind(&contact); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "JSON Error: " + err.Error(),
			})
		} else if result, err := services.SaveContact(db, contact); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Can't save record to database!",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Record has been saved!",
				"data":    result,
			})
		}
	}
}

func UpdateContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var contact models.Contact
		if err := ctx.ShouldBind(&contact); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			id, _ := strconv.Atoi(ctx.Query("id"))
			if result, err := services.UpdateContact(db, id, contact); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Can't save record to database!",
					"error":   err.Error(),
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

func DeleteContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		if result, err := services.DeleteContact(db, id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Can't deleted record to database!",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Record has been deleted!",
				"data":    result,
			})
		}
	}
}
