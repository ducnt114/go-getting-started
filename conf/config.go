package conf

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MySQL struct {
		Host            string `envconfig:"MYSQL_HOST"`
		Port            int64  `envconfig:"MYSQL_PORT"`
		User            string `envconfig:"MYSQL_USER"`
		Password        string `envconfig:"MYSQL_PASSWORD"`
		DB              string `envconfig:"MYSQL_DBNAME"`
		MigrationFolder string `envconfig:"MYSQL_MIGRATION_FOLDER"`
	}
	Sentry struct {
		Dsn string `envconfig:"SENTRY_DSN"`
	}
}

var GlobalConfig *Config

func InitConfig() error {
	GlobalConfig = &Config{}
	_ = godotenv.Load(".env")
	err := envconfig.Process("", GlobalConfig)
	return err
}
