package endpoints

import (
	"context"

	"github.com/alfuckk/fumin/internal/novel/service"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/fx"
)

type Endpoints struct {
	HelloEndpoint endpoint.Endpoint
}
type EndpointParams struct {
	fx.In
	Service service.Service
}

func New(params EndpointParams) Endpoints {
	return Endpoints{
		HelloEndpoint: makeHelloEndpoint(params.Service),
	}
}

func makeHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Hello(ctx)
	}
}
