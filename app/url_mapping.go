package app

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/bookstore-user-api/controller"
)

func mapUrls(router *gin.Engine) {
	router.GET("/ping", controller.Ping)

	router.GET("/users/:id", controller.GetUser)
	router.POST("/users", controller.CreateUser)
}
