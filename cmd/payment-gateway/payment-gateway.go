package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Payment Gateway")

	mux := http.NewServeMux()

	// Health check endpoint to check the liveness (health) of the service
	mux.HandleFunc("GET /healthz", livenessHandler)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("error occurred while trying to run server: %v", err)
	}
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
