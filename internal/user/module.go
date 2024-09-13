package user

import (
	"go.uber.org/fx"

	"github.com/alfuckk/fumin/internal/novel/endpoints"
	"github.com/alfuckk/fumin/internal/novel/service"
	"github.com/alfuckk/fumin/internal/novel/transport"
	"github.com/alfuckk/fumin/pkg/configfx"
	"github.com/alfuckk/fumin/pkg/dbfx"
	"github.com/alfuckk/fumin/pkg/logfx"
	"github.com/alfuckk/fumin/pkg/redisfx"
)

var Module = fx.Module("user",
	fx.Provide(
		configfx.New,
		logfx.NewDevelopmentLogger,
		redisfx.New,
		dbfx.New,
		service.New,
		endpoints.New,
		transport.NewHTTPHandler,
	),
	fx.Supply("user"),
)
