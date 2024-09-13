package dbfx

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	DSN string
}

type Database struct {
	*bun.DB
}

func New(p Params) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(p.DSN)))
	return bun.NewDB(sqldb, pgdialect.New())
}
