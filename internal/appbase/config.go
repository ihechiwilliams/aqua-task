package appbase

import (
	"time"

	"aqua-backend/pkg/config"
)

type Config struct {
	ServerPort                 string `env:"SERVER_PORT" env-default:"8080"`
	ServiceName                string `env:"SERVICE_NAME" env-default:"aqua-backend"`
	NotificationServerPort     string `env:"NOTIFICATION_SERVER_PORT" env-default:"9090"`
	NotificationGRPCServerPort string `env:"NOTIFICATION_GRPC_SERVER_PORT" env-default:"50051"`
	NotificationServiceName    string `env:"NOTIFICATION_SERVICE_NAME" env-default:"notification"`
	ServerTimeout              int64  `env:"SERVER_TIMEOUT" env-default:"120"`
	Env                        string `env:"ENV"`
	LogLevel                   string `env:"LOG_LEVEL" env-default:"debug"`

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
