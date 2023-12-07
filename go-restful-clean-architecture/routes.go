package main

import (
	"go-clean-architecture/src/modules/users/transport/gogin"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB, port int) {
	router := gin.Default()

	gogin.UsersRouteConfig(router, db)

	router.Run(":" + strconv.Itoa(port))
}
