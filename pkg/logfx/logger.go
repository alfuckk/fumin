package logfx

import (
	"os"

	"github.com/go-kit/log"
)

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
}

func Info(keyvals ...interface{}) {
	logger.Log(append([]interface{}{"level", "info"}, keyvals...)...)
}

func Error(keyvals ...interface{}) {
	logger.Log(append([]interface{}{"level", "error"}, keyvals...)...)
}

func New() log.Logger {
	return logger
}
