package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/alfuckk/fumin/internal/novel/endpoints"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

type handlerParams struct {
	fx.In
	Log      kitlog.Logger
	Endpoint endpoints.Endpoints
}

func NewHTTPHandler(p handlerParams) http.Handler {
	r := mux.NewRouter()
	r.Use(LoggingMiddleware(p.Log))

	r.Path("/hello").Handler(httptransport.NewServer(
		p.Endpoint.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	))

	return r
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
