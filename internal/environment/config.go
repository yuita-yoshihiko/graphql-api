package environment

var Conf Config

type (
	Config struct {
		Env            Env
		DBEnv          DBEnv
	}
	Env struct {
		IsLocal bool `envconfig:"ENV_IS_LOCAL" default:"false"`
	}
	DBEnv struct {
		DatabaseURL               string `envconfig:"DATABASE_URL" required:"true"`
		DatabasePassword          string `envconfig:"DATABASE_PASSWORD" required:"true"`
		SSLMode           string `envconfig:"DATABASE_SSL_MODE" required:"true"`
	}
)