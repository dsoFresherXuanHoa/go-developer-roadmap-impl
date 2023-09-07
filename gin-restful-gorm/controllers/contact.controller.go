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
		size, _ := strconv.Atoi(ctx.Query("size"))
		start, _ := strconv.Atoi(ctx.Query("start"))

		if contacts := services.FindAllContact(db, size, start); reflect.DeepEqual(contacts, models.Contacts{}) {
			ctx.JSON(http.StatusBadRequest, gin.H{
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
				"err": err.Error(),
			})
		} else {
			if contact := services.FindContactById(id, db); contact == (models.Contact{}) {
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
}

func SaveContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var contact models.Contact
		if err := ctx.ShouldBind(&contact); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			if !services.SaveContact(db, contact) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Can't save record to database!",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Record has been saved!",
				})
			}
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
			if !services.UpdateContact(db, id, contact) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Can't save record to database!",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Record has been saved!",
				})
			}
		}
	}
}

func DeleteContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		if !services.DeleteContact(db, id) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Can't deleted record to database!",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Record has been deleted!",
			})
		}
	}
}
