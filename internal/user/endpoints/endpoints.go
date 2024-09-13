package endpoints

import (
	"go.uber.org/fx"

	"github.com/alfuckk/fumin/internal/user/service"
)

type Endpoints struct{}

type EndpointParams struct {
	fx.In
	service service.Service
}

func New(ep EndpointParams) Endpoints {
	return Endpoints{}
}
