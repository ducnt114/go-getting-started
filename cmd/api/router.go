package api

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"go-getting-started/cmd/api/controller"
	"go-getting-started/middlewares"
)

func InitRouter(di *do.Injector) (*gin.Engine, error) {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middlewares.GenRequestId())
	r.Use(middlewares.GinZap())

	userController := controller.NewUserController(di)
	bookController := controller.NewBookController(di)

	v1 := r.Group("/api/v1")

	userGroup := v1.Group("/user")
	//userGroup.Use(middlewares.Auth(di))
	userGroup.GET("", userController.List)
	userGroup.GET("/:id", userController.GetUserById)
	userGroup.POST("", userController.Create)
	userGroup.PUT("/:id", userController.Update)
	userGroup.GET("/stream", userController.ServeSSE)

	//v1.DELETE("/user/:id", userController.DeleteUser)

	bookGroup := v1.Group("/book")
	//bookGroup.Use(middlewares.Auth(di))
	//bookGroup.Use(middlewares.Authz(di))
	//bookGroup.Use(middlewares.Authorization(di))
	bookGroup.GET("", bookController.List)
	bookGroup.GET("/data", bookController.GetData)

	return r, nil
}
