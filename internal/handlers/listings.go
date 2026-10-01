package handlers

import (
	"database/sql"
	"net/http"
)

// here we need to query the database so we need database connection from pool(from main need to pass in listing)
// dependency injection
// handler ko wrap krke return kiya aur wrapper function mein we receive the dependency
func List(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	}

}
