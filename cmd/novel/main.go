package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alfuckk/fumin/internal/novel"
	"github.com/alfuckk/fumin/pkg/configfx"
	"github.com/alfuckk/fumin/pkg/logfx"
	"github.com/oklog/oklog/pkg/group"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Invoke(startServer),
		novel.Module,
	).Run()
}

func startServer(lc fx.Lifecycle, handler http.Handler, logfx *logfx.Logger, cfg *configfx.Config) {
	port := fmt.Sprintf(":%d", cfg.GetInt("http.port"))
	srv := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logfx.Info("msg", fmt.Sprintf("Starting server on %s", port))
			var g group.Group
			g.Add(func() error {
				return srv.ListenAndServe()
			}, func(error) {
				logfx.Info("msg", "Stopping server")
			})

			go func() {
				if err := g.Run(); err != nil {
					logfx.Error("error", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logfx.Info("msg", "Stopping server")
			return srv.Shutdown(ctx)
		},
	})
}
