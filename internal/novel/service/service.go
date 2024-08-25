package service

import (
	"context"

	"github.com/go-kit/log"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Service interface {
	Hello(ctx context.Context) (string, error)
}

type service struct {
	cfg    *viper.Viper
	logger log.Logger
}
type ServiceParams struct {
	fx.In
	Config *viper.Viper
	Logger log.Logger
}

func New(p ServiceParams) Service {
	return &service{
		logger: p.Logger,
		cfg:    p.Config,
	}
}

func (s *service) Hello(ctx context.Context) (string, error) {
	return "Hello, World!", nil
}
