package router

import (
	"github.com/gin-gonic/gin"
	"go-getting-started/controller"
)

func InitRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	userController := &controller.UserController{}

	v1 := r.Group("/api/v1")
	v1.GET("/user", userController.ListUser)
	v1.GET("/user/:id", userController.GetUserDetail)
	v1.POST("/user", userController.CreateUser)
	v1.PUT("/user/:id", userController.UpdateUser)
	v1.DELETE("/user/:id", userController.DeleteUser)

	return r, nil
}
