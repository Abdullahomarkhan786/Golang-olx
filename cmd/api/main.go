package main

import (
	"Abdullahomarkhan786/Golang-olx/internal/config"
	"log"
	"net/http"
	"time"
)

func main() {

	cfg := config.MustLoad()

	//creates a new, independent router
	mux := http.NewServeMux() // create our own ServeMux instead of using the global DefaultServeMux

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json") // Set the header
		w.WriteHeader(http.StatusOK)                       //Send status + headers
		w.Write([]byte(`{"status":"OK"}`))                 //Send body

	})

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux, ////pass custom mux, mux (multiplexer) manages multiple handlers and routes and directs each incoming HTTP request to the appropriate handler.
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed %v", err)

	}
}
