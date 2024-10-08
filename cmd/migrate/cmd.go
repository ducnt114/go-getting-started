package migrate

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
	"go-getting-started/conf"
	"go.uber.org/zap"
)

var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate",
	Long:  `migrate`,
	Run: func(cmd *cobra.Command, args []string) {
		startMigration()
	},
}

func startMigration() {
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
		return
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return
	}
	msg := "migrate success"
	if errors.Is(err, migrate.ErrNoChange) {
		msg += ", but no changes"
	}
	zap.S().Info(msg)
	return
}
