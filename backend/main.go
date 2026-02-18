package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// Define a simple health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintln(w, "Health check passed!")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health check passed!"))
	})
	handler := enableCORS(mux)

	log.Println("starting server on :8080")

	http.ListenAndServe(":8080", handler)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
