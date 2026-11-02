package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Main application landing page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>🚀 Welcome to my Cloud-Native GitOps Application!</h1>\n")
		fmt.Fprintf(w, "<p>Status: Running smoothly inside Kubernetes.</p>\n")
	})

	// Health check endpoint for Kubernetes probes
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}