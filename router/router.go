package router

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-getting-started/controller"
	_ "go-getting-started/docs"
)

func InitRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("go-getting-started-api"),
		newrelic.ConfigLicense(""),
		newrelic.ConfigDistributedTracerEnabled(false),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		return nil, err
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userController := &controller.UserController{}

	v1 := r.Group("/api/v1")

	v1.Use(nrgin.Middleware(nrApp))

	//v1.GET("/user", userController.ListUser)
	v1.GET("/user/:id", userController.GetUserById)
	//v1.POST("/user", userController.CreateUser)
	//v1.PUT("/user/:id", userController.UpdateUser)
	//v1.DELETE("/user/:id", userController.DeleteUser)

	return r, nil
}
