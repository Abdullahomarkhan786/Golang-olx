package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Abdullahomarkhan786/olx-api/internal/config"
)

// . In a web server, a mux is basically the component that matches an incoming HTTP request to the correct handler
func main() {
	cfg := config.MustLoad()
	fmt.Println("Starting the olx server")

	mux := http.NewServeMux()
	//add route
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)       //flushes or sends the headers in the response(here we send the headers)
		w.Write([]byte(`{"status":"ok"}`)) //(here we send the data)writes data to connection as part of http reply and we use byte to convert normal string to byte

	})
	//initialising server with config
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 120,
	}
	log.Printf("Server is running on %s", srv.Addr)
	err := srv.ListenAndServe() //func ListenAndServe(addr string, handler Handler) error
	if err != nil {
		log.Fatalf("server failed %v", err)
	}

}
