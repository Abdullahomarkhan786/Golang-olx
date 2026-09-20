package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	//create our own router so only these routes which we create using it can be accessed
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //header is just a key value pair and whenever responese ends it sends all the headers
		w.WriteHeader(http.StatusOK)
		w.Write([]byte((`{"status":"ok"}`))) //sends the headers flushes to response
	})

	//initialise server with configurations
	srv := http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	err := srv.ListenAndServe() //we pass our custom config here
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
