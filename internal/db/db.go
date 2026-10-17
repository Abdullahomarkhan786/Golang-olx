package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("sql.Open:%w", err)

	}
	//• SetConnMaxIdleTime(d): Closes a connection if it has been sitting in the pool unused for the duration specified. If a connection is used frequently, its idle timer resets, and it stays open.
	//SetConnMaxLifetime(d): Closes a connection after the duration specified, measured strictly from the moment it was first created. It does not matter how busy or idle the connection has been.

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	//fail fast if database not ready then crash
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.ping:%w", err)
	}

	return db, nil

}
