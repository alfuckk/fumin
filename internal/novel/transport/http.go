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
	//
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	// 搜书
	r.Path("/api/novel/search").Handler(httptransport.NewServer(
		p.Endpoint.NovelListEndpoint,
		decodeNovelListRequest,
		encodeResponse,
	))
	// 获取书籍分类
	r.Path("/api/novel/category").Handler(httptransport.NewServer(
		p.Endpoint.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	))

	// 获取书籍详情
	r.Path("/api/chapter/detail").Handler(httptransport.NewServer(
		p.Endpoint.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	))

	// 书籍详情
	r.Path("/api/chapter/list").Handler(httptransport.NewServer(
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

func decodeNovelListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.NovelListReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
