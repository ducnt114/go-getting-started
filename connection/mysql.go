package connection

import (
	"fmt"
	"github.com/samber/do"
	"go-getting-started/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewMysqlConnection(di *do.Injector) (*gorm.DB, error) {
	cf := do.MustInvoke[*conf.Config](di)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		cf.MySQL.User, cf.MySQL.Password,
		cf.MySQL.Host, cf.MySQL.Port,
		cf.MySQL.DB,
	)
	dbOrm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
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
