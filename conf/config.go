package conf

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
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
		PublicKeyFilePath  string `envconfig:"JWT_PUBLIC_KEY_FILE_PATH"`
		PrivateKeyFilePath string `envconfig:"JWT_PRIVATE_KEY_FILE_PATH"`
	}

	GoogleAuth struct {
		ClientID     string `envconfig:"CLIENT_ID"`
		ClientSecret string `envconfig:"CLIENT_SECRET"`
		CallbackURL  string `envconfig:"CLIENT_CALLBACK_URL"`
	}

	Casbin struct {
		RBACModelPath  string `envconfig:"RBAC_MODEL_PATH"`
		RBACPolicyPath string `envconfig:"RBAC_POLICY_PATH"`
	}
}

func NewConfig(di *do.Injector) (*Config, error) {
	cf := &Config{}
	_ = godotenv.Load(".env")
	err := envconfig.Process("", cf)

	goth.UseProviders(
		google.New(
			cf.GoogleAuth.ClientID,
			cf.GoogleAuth.ClientSecret,
			cf.GoogleAuth.CallbackURL),
	)

	return cf, err
}
