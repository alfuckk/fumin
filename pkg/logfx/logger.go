package logfx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func (log Logger) Info(keyvals ...interface{}) {
	log.Infow("Info", append([]interface{}{"level", "info"}, keyvals...)...)
}

// Error 记录错误级别的日志
func (log Logger) Error(keyvals ...interface{}) {
	log.Errorw("Error", append([]interface{}{"level", "error"}, keyvals...)...)
}

func New() *Logger {
	config := zap.NewProductionConfig()

	// 配置日志的时间格式
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.TimeKey = "timestamp"

	// 构建日志对象
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &Logger{logger.Sugar()}
}

// NewDevelopmentLogger 创建一个用于开发环境的日志实例
func NewDevelopmentLogger() *Logger {
	config := zap.NewDevelopmentConfig()

	// 配置日志的时间格式
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.TimeKey = "timestamp"

	// 构建日志对象
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &Logger{logger.Sugar()}
}
