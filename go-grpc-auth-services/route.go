package main

import (
	"strconv"

	authGin "go-grpc-auth-services/src/module/transport/gogin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteConfig(db *gorm.DB, port int) {
	router := gin.Default()

	authGin.AuthRouteConfig(router, db)

	router.Run(":" + strconv.Itoa(port))
}
