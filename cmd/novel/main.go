package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"

	"github.com/alfuckk/fumin/internal/novel"
	"github.com/go-kit/log"
	"github.com/spf13/viper"
)

func main() {
	fx.New(
		novel.Module,
		fx.Invoke(startServer),
	).Run()
}

func startServer(lc fx.Lifecycle, handler http.Handler, log log.Logger, cfg *viper.Viper) {
	port := fmt.Sprintf(":%d", cfg.GetInt("http.port"))
	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Log("msg", fmt.Sprintf("Starting server on %s", port))
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
