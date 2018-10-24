package data

import (
	"context"
)

type Database interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Rows interface {
	Scan(dest ...interface{}) error
	Close() error
	Next() bool
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
