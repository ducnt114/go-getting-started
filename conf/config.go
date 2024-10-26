package conf

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/samber/do"
)

type Config struct {
	ApiService struct {
		Port int64 `envconfig:"API_PORT"`
	}

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

	JWT struct {
		PublicKeyFilePath string `envconfig:"JWT_PUBLIC_KEY_FILE_PATH"`
	}
}

func NewConfig(di *do.Injector) (*Config, error) {
	envConfig := &Config{}
	_ = godotenv.Load(".env")
	err := envconfig.Process("", envConfig)
	return envConfig, err
}
