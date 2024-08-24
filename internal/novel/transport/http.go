package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alfuckk/fumin/internal/novel/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Path("/hello").Handler(httptransport.NewServer(
		endpoints.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	))

	return r
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
