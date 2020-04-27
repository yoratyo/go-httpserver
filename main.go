package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main () {
	// Create route and handler
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8090",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server...")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

func indexHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go simple application!"))
}