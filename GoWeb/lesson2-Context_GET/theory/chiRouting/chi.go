package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, Users!")
		})
		r.Get("/{userID}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello User!"))
		})

		r.Get("/{userID}/posts/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello user posts!"))
		})
	})

	http.ListenAndServe(":3000", r)
}
