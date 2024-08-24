package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"

	"github.com/alfuckk/fumin/internal/novel/config"
	"github.com/alfuckk/fumin/internal/novel/endpoints"
	"github.com/alfuckk/fumin/internal/novel/service"
	"github.com/alfuckk/fumin/internal/novel/transport"
	"github.com/alfuckk/fumin/pkg/logger"
	"github.com/go-kit/log"
)

func main() {
	const app = "novel"
	fx.New(
		fx.Provide(
			config.New,
			logger.New,
			service.New,
			endpoints.New,
			transport.NewHTTPHandler,
		),
		fx.Invoke(startServer),
	).Run()
}

func startServer(lc fx.Lifecycle, handler http.Handler, log log.Logger) {
	server := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Log("msg", "Starting server on :8080")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Log("msg", "Stopping server")
			return server.Shutdown(ctx)
		},
	})

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Log("msg", "Shutting down server...")
}
