package services

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) (row *sqlx.Row)
	Rollback() error
	Commit() error
}
