package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)       //flushes or sends the headers in the response(here we send the headers)
	w.Write([]byte(`{"status":"ok"}`)) //(here we send the data)writes data to connection as part of http reply and we use byte to convert normal string to byte

}
