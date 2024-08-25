package logger

import (
	"os"

	"github.com/go-kit/log"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
}

func New(p Params) log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	return log.With(logger, "ts", log.DefaultTimestampUTC)
}
