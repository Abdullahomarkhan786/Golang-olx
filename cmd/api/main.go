package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatalf("server failed %v", err)

	}
}
