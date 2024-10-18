package repository

import (
	"context"
	"database/sql"
)

type database interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Store interface{}

type store struct {
	conn *sql.DB
	db   database
}

func NewStore(db *sql.DB) Store {
	return &store{
		conn: db,
		db:   db,
	}
}
