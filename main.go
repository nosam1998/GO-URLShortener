package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var DB *sql.DB

// TODO: Check for duplicate URLs and how to handle them.
// TODO: Make sure the shortUrl string is formatted properly

func main() {
	DB = ConnectToDB()

	if !DoesTableExist("url") {
		CreateUrlTable()
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		if err != nil {
			panic(err)
		}
	})

	router.HandleFunc("/{url}", RedirectRoute).Methods("GET")

	//router.HandleFunc("/api/shortener", shorten).Methods("GET", "POST")
	router.HandleFunc("/api/shorten", ShortenRoute).Methods("POST")

	//router.Use(loggingMiddleware)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
