package service

import (
	"context"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Service interface {
	Hello(ctx context.Context) (string, error)
	NovelList(ctx context.Context, keyword string, page int64) (string, error)
}

type service struct {
	cfg *viper.Viper
}
type ServiceParams struct {
	fx.In
	Config *viper.Viper
}

func New(p ServiceParams) Service {
	return &service{
		cfg: p.Config,
	}
}

func (s *service) Hello(ctx context.Context) (string, error) {
	return "Hello, World!", nil
}

func (s *service) NovelList(ctx context.Context, keyword string, page int64) (string, error) {
	return "Hello, World!", nil
}
