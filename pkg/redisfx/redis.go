package redisfx

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
}

type Redis struct {
	*redis.Client
}

func New(p Params) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
}
