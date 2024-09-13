package endpoints

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/fx"

	"github.com/alfuckk/fumin/internal/novel/service"
)

type Endpoints struct {
	HelloEndpoint     endpoint.Endpoint
	NovelListEndpoint endpoint.Endpoint
}
type EndpointParams struct {
	fx.In
	Service service.Service
}

func New(params EndpointParams) Endpoints {
	return Endpoints{
		HelloEndpoint:     makeHelloEndpoint(params.Service),
		NovelListEndpoint: makeNovelListEndpoint(params.Service),
	}
}

func makeHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Hello(ctx)
	}
}

func makeNovelListEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(NovelListReq)
		if !ok {
			return nil, fmt.Errorf("invalid request type")
		}
		return svc.NovelList(ctx, req.Keyword, req.Page)
	}
}
