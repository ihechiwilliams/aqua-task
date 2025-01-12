package appbase

import (
	"time"

	"aqua-backend/pkg/config"
)

type Config struct {
	ServerAddress             string `env:"SERVER_ADDRESS" env-default:"0.0.0.0:8080"`
	ServiceName               string `env:"SERVICE_NAME" env-default:"aqua-backend"`
	NotificationServerAddress string `env:"NOTIFICATION_SERVER_ADDRESS" env-default:"0.0.0.0:9090"`
	NotificationServiceName   string `env:"NOTIFICATION_SERVICE_NAME" env-default:"notification"`
	Port                      string `env:"PORT" env-default:"3000"`
	ServerTimeout             int64  `env:"SERVER_TIMEOUT" env-default:"120"`
	Env                       string `env:"ENV"`
	LogLevel                  string `env:"LOG_LEVEL" env-default:"debug"`
	SentryDSN                 string `env:"SENTRY_DSN"`

	// Database
	RabbitmqURL string `env:"RABBITMQ_URL" env-required:"true"`
	DatabaseURL string `env:"DATABASE_URL" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	c := new(Config)

	err := config.LoadConfig(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) HTTPServerTimeout() time.Duration {
	return time.Duration(c.ServerTimeout) * time.Second
}
