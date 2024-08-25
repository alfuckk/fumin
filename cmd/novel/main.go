package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alfuckk/fumin/internal/novel"
	"github.com/go-kit/log"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/viper"
	"go.uber.org/fx"
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

			var g group.Group
			g.Add(func() error {
				return http.ListenAndServe(port, handler)
			}, func(error) {
				log.Log("msg", "Stopping server")
			})

			go func() {
				if err := g.Run(); err != nil {
					log.Log("error", err)
				}
			}()

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
