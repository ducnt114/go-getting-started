package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/samber/do"
	"go.uber.org/zap"
)

func main() {
	err := startMigrateDB()
	if err != nil {
		panic(err)
	}
}

func startMigrateDB() error {
	di := do.New()
	defer func() {
		_ = di.Shutdown()
	}()

	config.Inject(di)
	cf := do.MustInvoke[*config.Config](di)
	databaseURL := fmt.Sprintf("mysql://%v:%v@tcp(%v:%v)/%v",
		cf.MySQL.User, cf.MySQL.Password,
		cf.MySQL.Host, cf.MySQL.Port,
		cf.MySQL.DB,
	)
	m, err := migrate.New(fmt.Sprintf("file://%v", cf.MySQL.MigrationFolder), databaseURL)
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
