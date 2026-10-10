package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // Set the header
	w.WriteHeader(http.StatusOK)                       //Send status + headers
	w.Write([]byte(`{"status":"OK"}`))                 //Send body

}
