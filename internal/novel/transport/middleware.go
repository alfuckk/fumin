package transport

import (
	"net/http"
	"time"

	kitlog "github.com/go-kit/log"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(logger kitlog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// 记录请求信息
			logger.Log("method", r.Method, "url", r.URL.Path, "time", start.Format(time.RFC3339))

			// 捕获响应
			rw := &responseWriter{w, http.StatusOK}
			next.ServeHTTP(rw, r)

			// 记录响应信息
			duration := time.Since(start)
			logger.Log("status", rw.statusCode, "duration", duration)
		})
	}
}
