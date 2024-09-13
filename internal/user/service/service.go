package service

import "go.uber.org/fx"

type (
	ServiceParam struct {
		fx.In
	}
	service struct{}
)

type Service interface{}

func New(sp ServiceParam) Service {
	return &service{}
}
