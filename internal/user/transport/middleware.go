package transport

import "net/http"

func LoggerMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
