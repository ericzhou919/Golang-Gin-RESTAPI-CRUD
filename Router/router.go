package router

import (
	"golang-gin/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", controller.GetUsers)
	router.GET("/user/:id", controller.GetUser)
	router.POST("/user", controller.Insert)
	router.DELETE("user/:id", controller.Delete)
	return router
}
