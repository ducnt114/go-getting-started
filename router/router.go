package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-getting-started/controller"
	_ "go-getting-started/docs"
	"go-getting-started/middleware"
	"go-getting-started/service"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	pingController := &controller.PingController{}
	userController := &controller.UserController{
		UserService: service.NewUserService(),
	}
	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("go-getting-started-api"),
		newrelic.ConfigLicense("5f86f1d81002b9f3695899e421a58fe0FFFFNRAL"),
		newrelic.ConfigDistributedTracerEnabled(false),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(nrgin.Middleware(nrApp))
	//r.Use(middleware.ResponseTime())

	r.GET("/ping", pingController.Ping)

	userGroup := r.Group("/user")
	userGroup.Use(middleware.ResponseTime())
	userGroup.GET("/by-gender", userController.GetUserByGender)
	userGroup.GET("", userController.ListUser)
	//userGroup.POST("", userController.CreateUser)
	userGroup.GET("/asdfdasf", handlerFunc)
	return r
}

func handlerFunc(ctx *gin.Context) {

}
