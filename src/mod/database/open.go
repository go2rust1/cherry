package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

var timeout = time.Second * 5

type client struct {
	db  *sqlx.DB
	err error
}

func open(driver, dsn string) (*sqlx.DB, error) {
	pipeline := make(chan client)
	go func() {
		_db, _err := sqlx.Connect(driver, dsn)
		pipeline <- client{db: _db, err: _err}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	select {
	case _client := <-pipeline:
		return _client.db, _client.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
