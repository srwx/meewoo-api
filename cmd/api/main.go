package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0" // version of meewoo application

type config struct {
	port int
	env  string
	db   struct {
		connectionString string
	}
}

type AppStatus struct {
	Status     string `json:"status"`
	Enviroment string `json:"enviroment"`
	Version    string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

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

func main() {
	var conf config

	flag.IntVar(&conf.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&conf.env, "env", "development", "Application enviroment (development | production)")
	flag.StringVar(&conf.db.connectionString, "dsn", "postgresql://postgres:Sorawong1@localhost:5432/go_movies?sslmode=disable", "Postgres connection string") // dsn: "Data Source Name"
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	/* Connection to db */
	db, err := openDB(conf)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := application{
		config: conf,
		logger: logger,
	}

	/* Create server */
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.port),
		Handler: app.routes(),
	}
	logger.Println("Server started on port", conf.port)

	/* Connect to server */
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
