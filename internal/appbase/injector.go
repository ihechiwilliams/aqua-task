package appbase

import (
	"github.com/gin-gonic/gin"
	"os"

	"aqua-backend/internal/api"
	"aqua-backend/internal/repositories/invoices"
	"aqua-backend/internal/repositories/invoicesitems"
	"aqua-backend/pkg/postgres"

	"gorm.io/gorm"

	v1 "aqua-backend/internal/api/v1"
	"aqua-backend/internal/repositories/activities"
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
	do.Provide(injector, func(i *do.Injector) (*v1.ActivitiesHandler, error) {
		return v1.NewActivitiesHandler(
			do.MustInvoke[*activities.SQLRepository](i),
		), nil
	})

	do.Provide(injector, func(i *do.Injector) (*v1.CustomersHandler, error) {
		return v1.NewCustomersHandler(
			do.MustInvoke[*customers.SQLRepository](i),
		), nil
	})

	do.Provide(injector, func(i *do.Injector) (*v1.InvoiceHandler, error) {
		return v1.NewInvoiceHandler(
			do.MustInvoke[*invoices.SQLRepository](i),
			do.MustInvoke[*invoicesitems.SQLRepository](i),
		), nil
	})

	do.Provide(injector, func(i *do.Injector) (*v1.API, error) {
		activitiesHandler := do.MustInvoke[*v1.ActivitiesHandler](i)
		customersHandler := do.MustInvoke[*v1.CustomersHandler](i)
		invoiceHandler := do.MustInvoke[*v1.InvoiceHandler](i)

		return v1.NewAPI(activitiesHandler, customersHandler, invoiceHandler), nil
	})

	do.Provide(injector, func(i *do.Injector) (*api.Routes, error) {
		v1API := do.MustInvoke[*v1.API](i)

		return api.NewRoutes(v1API), nil
	})

	// ===========================
	//	Database Config & Repo
	// ===========================
	do.Provide(injector, func(i *do.Injector) (*activities.SQLRepository, error) {
		gormDB := do.MustInvokeNamed[*gorm.DB](i, InjectorDatabase)
		return activities.NewSQLRepository(gormDB), nil
	})

	do.Provide(injector, func(i *do.Injector) (*customers.SQLRepository, error) {
		gormDB := do.MustInvokeNamed[*gorm.DB](i, InjectorDatabase)
		return customers.NewSQLRepository(gormDB), nil
	})

	do.Provide(injector, func(i *do.Injector) (*invoices.SQLRepository, error) {
		gormDB := do.MustInvokeNamed[*gorm.DB](i, InjectorDatabase)
		return invoices.NewSQLRepository(gormDB), nil
	})

	do.Provide(injector, func(i *do.Injector) (*invoicesitems.SQLRepository, error) {
		gormDB := do.MustInvokeNamed[*gorm.DB](i, InjectorDatabase)
		return invoicesitems.NewSQLRepository(gormDB), nil
	})

	do.ProvideNamed(injector, InjectorDatabase, func(i *do.Injector) (*gorm.DB, error) {
		return postgres.InitDB(
			serviceName, &postgres.Config{
				Name:            cfg.DatabaseName,
				Password:        cfg.DatabasePassword,
				PrimaryHost:     cfg.DatabasePrimaryHost,
				ReadReplicaHost: cfg.DatabaseReadReplicaHost,
				User:            cfg.DatabaseUsername,
				Port:            cfg.DatabasePort,
			},
		)
	})

	return injector
}
