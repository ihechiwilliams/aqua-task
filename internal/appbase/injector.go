package appbase

import (
	"aqua-backend/internal/repositories/notification"
	"aqua-backend/internal/repositories/resources"
	"aqua-backend/pkg/rabbitmq"
	"github.com/gin-gonic/gin"
	"os"

	"aqua-backend/internal/api"
	"aqua-backend/pkg/postgres"

	"gorm.io/gorm"

	v1 "aqua-backend/internal/api/v1"
	"aqua-backend/internal/repositories/customers"

	"github.com/rs/zerolog"
	"github.com/samber/do"
)

func NewInjector(serviceName string, cfg *Config) *do.Injector {
	injector := do.New()

	// ===========================
	//	Service Configs (logging, open-api,...)
	// ===========================
	do.Provide(injector, func(i *do.Injector) (*zerolog.Logger, error) {
		logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
		if err != nil {
			return nil, err
		}

		logger := zerolog.New(os.Stdout).
			Level(logLevel).
			With().
			Str("serviceName", serviceName).
			Logger()

		return &logger, nil
	})

	do.ProvideNamed(injector, InjectorApplicationRouter, func(i *do.Injector) (*gin.Engine, error) {

		return NewRouterGin(serviceName, cfg.HTTPServerTimeout()), nil
	})

	// ===========================
	//	API services & Routes
	// ===========================
	do.Provide(injector, func(i *do.Injector) (*v1.CustomersHandler, error) {
		return v1.NewCustomersHandler(
			do.MustInvoke[*customers.SQLRepository](i),
		), nil
	})

	do.Provide(injector, func(i *do.Injector) (*v1.ResourcesHandler, error) {
		return v1.NewResourcesHandler(
			do.MustInvoke[*resources.SQLRepository](i),
			do.MustInvoke[*customers.SQLRepository](i),
		), nil
	})

	do.Provide(injector, func(i *do.Injector) (*v1.NotificationsHandler, error) {
		return v1.NewNotificationHandler(
			do.MustInvoke[*notification.SQLRepository](i),
		), nil
	})

	do.Provide(injector, func(i *do.Injector) (*v1.API, error) {
		customersHandler := do.MustInvoke[*v1.CustomersHandler](i)
		resourcesHandler := do.MustInvoke[*v1.ResourcesHandler](i)
		notificationHandler := do.MustInvoke[*v1.NotificationsHandler](i)

		return v1.NewAPI(customersHandler, resourcesHandler, notificationHandler), nil
	})

	do.Provide(injector, func(i *do.Injector) (*api.Routes, error) {
		v1API := do.MustInvoke[*v1.API](i)

		return api.NewRoutes(v1API), nil
	})

	// ===========================
	//	Database Config & Repo
	// ===========================
	do.Provide(injector, func(i *do.Injector) (*customers.SQLRepository, error) {
		gormDB := do.MustInvokeNamed[*gorm.DB](i, InjectorDatabase)
		return customers.NewSQLRepository(gormDB), nil
	})

	do.Provide(injector, func(i *do.Injector) (*resources.SQLRepository, error) {
		gormDB := do.MustInvokeNamed[*gorm.DB](i, InjectorDatabase)
		return resources.NewSQLRepository(gormDB), nil
	})

	do.Provide(injector, func(i *do.Injector) (*notification.SQLRepository, error) {
		gormDB := do.MustInvokeNamed[*gorm.DB](i, InjectorDatabase)
		return notification.NewSQLRepository(gormDB), nil
	})

	do.ProvideNamed(injector, InjectorDatabase, func(i *do.Injector) (*gorm.DB, error) {
		return postgres.InitDB(
			cfg.DatabaseURL,
		)
	})

	do.ProvideNamed(injector, InjectorRabbitmq, func(i *do.Injector) (*rabbitmq.RabbitMQ, error) {
		return rabbitmq.NewRabbitMQ(cfg.RabbitmqURL)
	})

	return injector
}
