package service

import (
	"context"

	"github.com/alfuckk/fumin/pkg/configfx"
	"github.com/alfuckk/fumin/pkg/logfx"
	"go.uber.org/fx"
)

type Service interface {
	Hello(ctx context.Context) (string, error)
	NovelList(ctx context.Context, keyword string, page int64) (string, error)
}

type service struct {
	cfg *configfx.Config
	log *logfx.Logger
}
type ServiceParams struct {
	fx.In
	Config *configfx.Config
	Log    *logfx.Logger
}

func New(p ServiceParams) Service {
	return &service{
		cfg: p.Config,
		log: p.Log,
	}
}

func (s *service) Hello(ctx context.Context) (string, error) {
	data := fetchNovel()
	// categories := fetchNovelCategory(data)
	s.log.Info("status", data)
	return "Hello, World!", nil
}

func (s *service) NovelList(ctx context.Context, keyword string, page int64) (string, error) {
	return "Hello, World!", nil
}
