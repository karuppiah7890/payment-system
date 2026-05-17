package main

import (
	"fmt"
	"log"
	"net/http"
	"payment/internal"
)

func main() {
	fmt.Println("Payment Processor")

	mux := http.NewServeMux()

	internal.AddHealthCheckEndpoint(mux)

	internal.AddMetricsEndpoint(mux)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("error occurred while trying to run server: %v", err)
	}
}
