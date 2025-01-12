package main

import (
	"context"
	"fmt"
	"net/http"

	"aqua-backend/cmd"
	"aqua-backend/internal/appbase"
	"aqua-backend/pkg/rabbitmq"
	"aqua-backend/pkg/signals"

	"github.com/rs/zerolog/log"
	"github.com/samber/do"
)

const (
	serviceName = "aqua-backend.server"
)

func main() {
	cmd.Execute()

	ctx, mainCtxStop := context.WithCancel(context.Background())

	app := appbase.New(
		appbase.Init(serviceName),
		appbase.WithDependencyInjector(),
	)
	defer app.Shutdown()
	fmt.Println(serviceName)

	// Initialize dependencies
	rmq := do.MustInvokeNamed[*rabbitmq.RabbitMQ](app.Injector, appbase.InjectorRabbitmq)

	// Simulate publishing a message
	queue := "notifications_queue"

	err := rmq.DeclareQueue(queue)
	if err != nil {
		log.Error().Msgf("Failed to declare queue: %v", err)
	}

	message := `{"user_id": "123", "message": "New resource created"}`

	err = rmq.PublishMessage(queue, message)
	if err != nil {
		log.Error().Msgf("Failed to publish message: %v", err)
	}

	log.Info().Msg("Message published!")

	// HTTP server setup
	router := buildRouter(app)
	httpServer := &http.Server{
		Addr:              app.Config.ServerAddress,
		Handler:           router,
		ReadHeaderTimeout: app.Config.HTTPServerTimeout(),
	}

	// Graceful shutdown for HTTP server
	signals.HandleSignals(ctx, mainCtxStop, func() {
		shutdownErr := httpServer.Shutdown(ctx)
		if shutdownErr != nil {
			log.Fatal().Err(shutdownErr).Msg("HTTP server shutdown failed")
		}
	})

	// Start HTTP server
	log.Info().Msgf("Started HTTP server on %s", app.Config.ServerAddress)

	// Start HTTP server in the main goroutine
	serverErr := httpServer.ListenAndServe()
	if serverErr != nil {
		log.Err(serverErr).Msg("HTTP server stopped")
	}

	// Wait for cancellation signal
	<-ctx.Done()
}
