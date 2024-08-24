package logger

import (
	"os"

	"github.com/go-kit/log"
)

func New() log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	return log.With(logger, "ts", log.DefaultTimestampUTC)
}
