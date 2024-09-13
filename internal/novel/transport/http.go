package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/fx"

	"github.com/alfuckk/fumin/internal/novel/endpoints"
	"github.com/alfuckk/fumin/pkg/logfx"
)

type handlerParams struct {
	fx.In
	Log      *logfx.Logger
	Endpoint endpoints.Endpoints
}

func NewHTTPHandler(p handlerParams) http.Handler {
	r := chi.NewRouter()
	r.Use(LoggingMiddleware(p.Log))
	//
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	searchNovel := httptransport.NewServer(
		p.Endpoint.NovelListEndpoint,
		decodeNovelListRequest,
		encodeResponse,
	)
	r.Method(http.MethodGet, "/api/novel/search", searchNovel)
	// 搜书
	categoryNovel := httptransport.NewServer(
		p.Endpoint.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	)
	r.Method(http.MethodGet, "/api/novel/category", categoryNovel)
	// 获取书籍分类
	detailNovel := httptransport.NewServer(
		p.Endpoint.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	)
	r.Method(http.MethodGet, "/api/chapter/detail", detailNovel)
	// 获取书籍详情
	ListChapter := httptransport.NewServer(
		p.Endpoint.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	)
	r.Method(http.MethodGet, "/api/chapter/list", ListChapter)
	// 书籍详情

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
