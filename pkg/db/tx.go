package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Commit() error
	Rollback() error
	// Add any additional methods you need here
}

type CustomTx struct {
	*sqlx.Tx
}

func (tx *CustomTx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *CustomTx) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx *CustomTx) Get(dest interface{}, query string, args ...interface{}) error {
	return tx.Tx.Get(dest, query, args...)
}
