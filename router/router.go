package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-getting-started/controller"
	_ "go-getting-started/docs"
)

func InitRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userController := &controller.UserController{}

	v1 := r.Group("/api/v1")
	//v1.GET("/user", userController.ListUser)
	v1.GET("/user/:id", userController.GetUserById)
	//v1.POST("/user", userController.CreateUser)
	//v1.PUT("/user/:id", userController.UpdateUser)
	//v1.DELETE("/user/:id", userController.DeleteUser)

	return r, nil
}
