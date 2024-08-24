package novel

import (
	"github.com/alfuckk/fumin/internal/novel/config"
	"github.com/alfuckk/fumin/internal/novel/endpoints"
	"github.com/alfuckk/fumin/internal/novel/service"
	"github.com/alfuckk/fumin/internal/novel/transport"
	"github.com/alfuckk/fumin/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Module("novel",
	fx.Provide(
		config.New,
		logger.New,
		service.New,
		endpoints.New,
		transport.NewHTTPHandler,
	),
)
