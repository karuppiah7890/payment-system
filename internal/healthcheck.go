package internal

import "net/http"

func AddHealthCheckEndpoint(mux *http.ServeMux) {
	// Health check endpoint to check the liveness (health) of the service
	mux.HandleFunc("GET /healthz", livenessHandler)
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
