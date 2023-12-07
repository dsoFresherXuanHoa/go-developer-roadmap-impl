package gogin

import (
	"go-grpc-auth-services/src/module/business"
	"go-grpc-auth-services/src/module/entity"
	"go-grpc-auth-services/src/module/storage"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var register entity.RegisterCreatable
		if err := ctx.ShouldBind(&register); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":    nil,
				"message": "Can't parse user request to struct",
				"error":   err.Error(),
			})
		} else {
			storage := storage.NewSQLStore(db)
			business := business.NewRegisterBusiness(storage)
			if _, _, err := business.SignUp(ctx.Request.Context(), &register, db); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"data":    nil,
					"message": "Can't save a user, account to database",
					"error":   err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data":    true,
					"message": "Success",
					"error":   nil,
				})
			}
		}
	}
}
