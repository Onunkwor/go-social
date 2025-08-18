package db

import (
	"context"
	"database/sql"
	"time"
)

// Internal packages should not know about the outside world, so we don't import "github/onunkwor/social/internal/env" here and also the dbconfig struct.
// New creates a new database connection with the provided configuration.
func New(dsn string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	idleTime, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(idleTime)
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
