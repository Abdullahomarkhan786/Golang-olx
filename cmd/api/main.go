package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Abdullahomarkhan786/olx-api/internal/config"
	"github.com/Abdullahomarkhan786/olx-api/internal/db"
	"github.com/Abdullahomarkhan786/olx-api/internal/handlers"
)

// . In a web server, a mux is basically the component that matches an incoming HTTP request to the correct handler
func main() {
	cfg := config.MustLoad()
	_, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect:%v", err) //to print struct

	}
	fmt.Println("Starting the olx server")

	mux := http.NewServeMux()
	//add route
	mux.HandleFunc("GET /healthz", handlers.Health)
	//initialising server with config
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 120,
	}
	log.Printf("Server is running on %s", srv.Addr)
	err = srv.ListenAndServe() //func ListenAndServe(addr string, handler Handler) error
	if err != nil {
		log.Fatalf("server failed %v", err)
	}

}
