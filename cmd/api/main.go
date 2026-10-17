package main

import (
	"Abdullahomarkhan786/Golang-olx/internal/config"
	"Abdullahomarkhan786/Golang-olx/internal/db"
	"Abdullahomarkhan786/Golang-olx/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	cfg := config.MustLoad()
	_, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect:%v", err)
	}

	fmt.Println("database connection ")
	fmt.Println("starting olx server")
	//creates a new, independent router
	mux := http.NewServeMux() // create our own ServeMux instead of using the global DefaultServeMux

	mux.HandleFunc("GET /healthz", handlers.Health)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux, ////pass custom mux, mux (multiplexer) manages multiple handlers and routes and directs each incoming HTTP request to the appropriate handler.
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed %v", err)

	}
}
