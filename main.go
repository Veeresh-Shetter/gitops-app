package main

import (
	"fmt"
	"net/http"
)

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Main application landing page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>🚀 Welcome to my Cloud-Native GitOps Application!</h1>")
		fmt.Fprintln(w, "<p>Status: Running smoothly inside Kubernetes.</p>")
	})

	// Health check endpoint for Kubernetes probes
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	return mux
}

func main() {
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", setupRouter()); err != nil {
		panic(err)
	}
}
