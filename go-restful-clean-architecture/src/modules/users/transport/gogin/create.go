package gogin

import (
	"go-clean-architecture/src/modules/users/business"
	"go-clean-architecture/src/modules/users/entity"
	"go-clean-architecture/src/modules/users/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.UserCreatable
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":    nil,
				"message": "Can't parse user request to struct",
				"error":   err,
			})
		} else {
			storage := storage.NewSQLStore(db)
			business := business.NewCreateBusiness(storage)
			if err := business.Create(ctx.Request.Context(), &user); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"data":    nil,
					"message": "Can't save a user to database",
					"error":   err,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data":    user.Username,
					"message": "Success",
					"error":   nil,
				})
			}
		}
	}
}
