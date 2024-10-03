package router

import (
	"github.com/gin-gonic/gin"
	"go-getting-started/controller"
	"go-getting-started/service"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	pingController := &controller.PingController{}
	userController := &controller.UserController{
		UserService: service.NewUserService(),
	}

	r.GET("/ping", pingController.Ping)

	userGroup := r.Group("/user")
	userGroup.GET("/by-gender", userController.GetUserByGender)
	userGroup.GET("", userController.ListUser)
	//userGroup.POST("", userController.CreateUser)
	userGroup.GET("/asdfdasf", handlerFunc)
	return r
}

func handlerFunc(ctx *gin.Context) {

}
