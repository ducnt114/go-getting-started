package router

import (
	"github.com/gin-gonic/gin"
	"go-getting-started/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	pingController := &controller.PingController{}
	userController := &controller.UserController{}

	r.GET("/ping", pingController.Ping)

	userGroup := r.Group("/user")
	userGroup.GET("/:id", userController.GetUserById)
	userGroup.GET("", userController.ListUser)
	//userGroup.POST("", userController.CreateUser)

	return r
}
