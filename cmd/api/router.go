package api

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-getting-started/cmd/api/controller"
	_ "go-getting-started/docs"
	"go-getting-started/middlewares"
)

func InitRouter(di *do.Injector) (*gin.Engine, error) {
	//cf := do.MustInvoke[*conf.Config](di)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middlewares.GenRequestId())
	r.Use(middlewares.GinZap())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userController := controller.NewUserController(di)

	v1 := r.Group("/api/v1")

	v1.GET("/user", userController.List)
	v1.GET("/user/:id", userController.GetUserById)
	v1.POST("/user", userController.Create)
	v1.PUT("/user/:id", userController.Update)
	//v1.DELETE("/user/:id", userController.DeleteUser)

	return r, nil
}
