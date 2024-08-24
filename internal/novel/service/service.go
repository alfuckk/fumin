package service

import (
	"context"

	"github.com/go-kit/log"
)

type Service interface {
	Hello(ctx context.Context) (string, error)
}

type service struct {
	logger log.Logger
}

func New(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) Hello(ctx context.Context) (string, error) {

	return "Hello, World!", nil
}
