package service

import (
	"context"

	"go.uber.org/fx"

	"github.com/alfuckk/fumin/pkg/configfx"
	"github.com/alfuckk/fumin/pkg/dbfx"
	"github.com/alfuckk/fumin/pkg/logfx"
	"github.com/alfuckk/fumin/pkg/redisfx"
)

type Service interface {
	Hello(ctx context.Context) (string, error)
	NovelList(ctx context.Context, keyword string, page int64) (string, error)
}

type service struct {
	cfg   *configfx.Config
	log   *logfx.Logger
	db    *dbfx.Database
	redis *redisfx.Redis
}

type ServiceParams struct {
	fx.In
	Config *configfx.Config
	Log    *logfx.Logger
	DB     *dbfx.Database
	Redis  *redisfx.Redis
}

func New(p ServiceParams) Service {
	return &service{
		cfg:   p.Config,
		log:   p.Log,
		db:    p.DB,
		redis: p.Redis,
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
