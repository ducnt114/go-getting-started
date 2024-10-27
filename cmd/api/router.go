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
	authController := controller.NewAuthController(di)

	v1 := r.Group("/api/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/login", authController.PasswordLogin)

	userGroup := v1.Group("/user")
	userGroup.Use(middlewares.Auth(di))
	userGroup.GET("", userController.List)
	userGroup.GET("/:id", userController.GetUserById)
	userGroup.POST("", userController.Create)
	userGroup.PUT("/:id", userController.Update)
	//v1.DELETE("/user/:id", userController.DeleteUser)

	return r, nil
}
