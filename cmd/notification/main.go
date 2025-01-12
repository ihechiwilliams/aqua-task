package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"aqua-backend/internal/api/v1"
	"aqua-backend/internal/appbase"
	"aqua-backend/internal/notificationconsumer"
	"aqua-backend/internal/repositories/notification"
	"aqua-backend/pkg/rabbitmq"
	"aqua-backend/pkg/signals"
	"aqua-backend/proto"

	"github.com/rs/zerolog/log"
	"github.com/samber/do"
	"google.golang.org/grpc"
)

const (
	serviceName = "notification.server"
)

func main() {
	ctx, mainCtxStop := context.WithCancel(context.Background())

	app := appbase.New(
		appbase.Init(serviceName),
		appbase.WithDependencyInjector(),
	)
	defer app.Shutdown()
	fmt.Println(serviceName)

	// Initialize dependencies
	rmq := do.MustInvokeNamed[*rabbitmq.RabbitMQ](app.Injector, appbase.InjectorRabbitmq)
	notificationRepo := do.MustInvoke[*notification.SQLRepository](app.Injector)

	// Start Notification Consumer in a separate goroutine
	go func() {
		queueName := "notifications_queue"
		err := rmq.DeclareQueue(queueName)
		if err != nil {
			log.Fatal().Msgf("Failed to declare queue: %v", err)
		}

		log.Info().Msg("Starting notification consumer...")
		notificationconsumer.StartNotificationConsumer(rmq, queueName, notificationRepo)
	}()

	// Start gRPC server
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal().Msgf("Failed to listen on port 50051: %v", err)
		}

		grpcServer := grpc.NewServer()
		proto.RegisterNotificationServiceServer(grpcServer, &v1.NotificationServer{
			Repo: notificationRepo,
		})

		log.Info().Msg("gRPC server is running on port 50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal().Msgf("Failed to serve gRPC server: %v", err)
		}
	}()

	// HTTP server setup
	router := buildRouter(app)
	httpServer := &http.Server{
		Addr:              app.Config.NotificationServerAddress,
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
	log.Info().Msgf("Started HTTP server on %s", app.Config.NotificationServerAddress)

	// Start HTTP server in the main goroutine
	serverErr := httpServer.ListenAndServe()
	if serverErr != nil {
		log.Err(serverErr).Msg("HTTP server stopped")
	}

	// Wait for cancellation signal
	<-ctx.Done()
}
