package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseUrl string) (*sql.DB, error) {
	//open the connection
	db, err := sql.Open("pgx", databaseUrl) //give (driver name,source name) and it returns sql db pointer and error
	if err != nil {
		return nil, fmt.Errorf("sql.Open:%w", err) //db=nill since we got error so we return the formatted error
	}
	//configure this connection pool
	db.SetMaxOpenConns(25) //SetMaxOpenConns sets the maximum number of open connections to the database.,the default is 0 which means for each query we will have a new connection
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	//fail fast and it receives a context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //pass a parent context
	defer cancel()
	err = db.PingContext(ctx) //we pass timeout here and we try to ping the db if it doesn't connect in 5 sec we get the error
	if err != nil {
		return nil, fmt.Errorf("db.ping: %w", err)
	}
	return db, nil //no error

}
