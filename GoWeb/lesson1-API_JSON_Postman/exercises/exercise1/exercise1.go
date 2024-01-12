package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	// Create a new router
	rt := chi.NewRouter()

	// Endpoint
	rt.Get("/ping", ping)

	// Start server
	if err := http.ListenAndServe(":3000", rt); err != nil {
		panic(err)
	}

}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong!"))
}
