package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	// Create a new router
	rt := chi.NewRouter()

	// Endpoint
	rt.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		//set code
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!!!"))

	})

	// Start server
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}

}
