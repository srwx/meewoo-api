package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0" // version of meewoo application

type config struct {
	port int
	env  string
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

func main() {
	var conf config

	flag.IntVar(&conf.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&conf.env, "env", "development", "Application enviroment (development | production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: conf,
		logger: logger,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.port),
		Handler: app.routes(),
	}

	logger.Println("Server started on port", conf.port)

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
