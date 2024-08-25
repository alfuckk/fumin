package novel

import (
	"github.com/alfuckk/fumin/internal/novel/endpoints"
	"github.com/alfuckk/fumin/internal/novel/service"
	"github.com/alfuckk/fumin/internal/novel/transport"
	"github.com/alfuckk/fumin/pkg/config"
	"github.com/alfuckk/fumin/pkg/logfx"
	"go.uber.org/fx"
)

var Module = fx.Module("novel",
	fx.Provide(
		config.New,
		logfx.New,
		service.New,
		endpoints.New,
		transport.NewHTTPHandler,
	),
	fx.Supply("novel"),
)
