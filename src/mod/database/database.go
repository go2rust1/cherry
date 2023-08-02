package database

import (
	"github.com/jmoiron/sqlx"
)

type Database interface {
	SetDSN(string)
	SetSchema(string)
	SetTable(string)
	Bind(interface{})
	DSN() string
	DST() string
	Tag() []string
	Open() (*sqlx.DB, error)
	Share(*sqlx.DB)
	Insert(...interface{})
}
