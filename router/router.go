package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-getting-started/conf"
	"go-getting-started/controller"
	_ "go-getting-started/docs"
	"go-getting-started/model"
	"go-getting-started/repository"
	"go-getting-started/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitRouter() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dbOrm, err := initDbConnection()
	if err != nil {
		return nil, err
	}
	err = dbOrm.AutoMigrate(&model.User{}, &model.Book{})
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(dbOrm)
	userService := service.NewUserService(userRepo)

	userController := &controller.UserController{
		UserService: userService,
	}

	v1 := r.Group("/api/v1")

	//v1.GET("/user", userController.ListUser)
	v1.GET("/user/:id", userController.GetUserById)
	v1.POST("/user", userController.Create)
	//v1.PUT("/user/:id", userController.UpdateUser)
	//v1.DELETE("/user/:id", userController.DeleteUser)

	return r, nil
}

func initDbConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		conf.GlobalConfig.MySQL.User, conf.GlobalConfig.MySQL.Password,
		conf.GlobalConfig.MySQL.Host, conf.GlobalConfig.MySQL.Port,
		conf.GlobalConfig.MySQL.DB,
	)
	dbOrm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	db, err := dbOrm.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	return dbOrm, nil
}
