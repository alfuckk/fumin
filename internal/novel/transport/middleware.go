package transport

import (
	"net/http"
	"time"

	"github.com/alfuckk/fumin/pkg/logfx"
)

func LoggingMiddleware(logger *logfx.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// 记录请求信息
			logger.Info("method", r.Method, "url", r.URL.Path, "time", start.Format(time.RFC3339))

			// 捕获响应
			rw := &responseWriter{w, http.StatusOK}
			next.ServeHTTP(rw, r)

			// 记录响应信息
			duration := time.Since(start)
			logger.Info("status", rw.statusCode, "duration", duration)
		})
	}
}
