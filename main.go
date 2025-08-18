package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening on port 8080...")
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed to start: %e", err)
	}
}
