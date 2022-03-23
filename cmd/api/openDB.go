package main

import (
	"context"
	"database/sql"
	"time"
)

/*
** *sql.DB is connection pool (use to connect to db)
 */
func openDB(conf config) (*sql.DB, error) {
	db, err := sql.Open("postgres", conf.db.connectionString)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	/* PingContext verifies the connection to the database is still alive. */
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
