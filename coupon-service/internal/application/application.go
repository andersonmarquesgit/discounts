package application

import (
	"coupon/internal/config"
	"coupon/internal/infrastructure/logging"
	"coupon/internal/interfaces/http/handlers"
	"coupon/internal/usecases"
	"github.com/gocql/gocql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"time"
)

type Application struct {
	Config      *config.Config
	Cassandra   *gocql.Session
	UseCases    *usecases.UseCases
	Handlers    *handlers.Handlers
	NewRelicApp *newrelic.Application
}

func NewApplication(cfg *config.Config) *Application {
	var newRelicApp *newrelic.Application
	var err error

	if cfg.NewRelic.Enabled && cfg.NewRelic.LicenseKey != "" {
		newRelicApp, err = newrelic.NewApplication(
			newrelic.ConfigAppName("chargeback-api"),
			newrelic.ConfigLicense(cfg.NewRelic.LicenseKey),
			newrelic.ConfigDistributedTracerEnabled(true),
			newrelic.ConfigAppLogForwardingEnabled(true),
		)
		if err != nil {
			log.Fatalf("Failed to create New Relic application: %v", err)
		}

		logging.InitializeLogger(newRelicApp)
		logging.Logger.Info("Application started with New Relic integration")

	} else {
		logging.InitializeLogger(nil)
		logging.Logger.Info("New Relic is disabled. Starting without New Relic integration")
	}

	// Initialize the global logger
	logging.InitializeLogger(newRelicApp)
	logging.Logger.Info("Application started with New Relic integration")

	// Connect to database

	// Connect to RabbitMQ

	// Initialize producers

	// Initialize repositories

	// Initialize use cases
	couponGenerateUseCase := usecases.NewCouponGenerateUseCase()
	useCases := usecases.NewUseCases(couponGenerateUseCase)

	// Initialize specific handlers
	couponGenerateHandler := handlers.NewCouponGenerateHandler(useCases.CouponGenerateUseCase)

	// Initialize handlers
	handlers := handlers.NewHandlers(couponGenerateHandler)

	return &Application{
		Config:      cfg,
		UseCases:    useCases,
		Handlers:    handlers,
		NewRelicApp: newRelicApp,
	}
}

func (app *Application) Close() {
	if app.Cassandra != nil {
		app.Cassandra.Close()
	}

	if app.NewRelicApp != nil {
		app.NewRelicApp.Shutdown(10 * time.Second)
	}
}
