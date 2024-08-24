package endpoints

import (
	"context"

	"github.com/alfuckk/fumin/internal/novel/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	HelloEndpoint endpoint.Endpoint
}

func New(svc service.Service) Endpoints {
	return Endpoints{
		HelloEndpoint: makeHelloEndpoint(svc),
	}
}

func makeHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Hello(ctx)
	}
}
