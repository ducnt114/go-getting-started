package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-getting-started/conf"
	"go.uber.org/zap"
)

func init() {
	conf.InitConfig()
}

func main() {
	err := startMigrateDB()
	if err != nil {
		panic(err)
	}
}

func startMigrateDB() error {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}

	databaseURL := fmt.Sprintf("mysql://%v:%v@tcp(%v:%v)/%v",
		conf.GlobalConfig.MySQL.User, conf.GlobalConfig.MySQL.Password,
		conf.GlobalConfig.MySQL.Host, conf.GlobalConfig.MySQL.Port,
		conf.GlobalConfig.MySQL.DB,
	)
	m, err := migrate.New(fmt.Sprintf("file://%v", conf.GlobalConfig.MySQL.MigrationFolder), databaseURL)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	msg := "migrate success"
	if errors.Is(err, migrate.ErrNoChange) {
		msg += ", but no changes"
	}
	zap.S().Info(msg)
	return nil
}
