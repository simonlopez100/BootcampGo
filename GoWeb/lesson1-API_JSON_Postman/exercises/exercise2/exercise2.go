package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/go-chi/chi"
)

type PersonToGreet struct {
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
}

func greetingsHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	fmt.Println(r.Body)

	var personToGreet PersonToGreet
	err := decoder.Decode(&personToGreet)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}
	fmt.Printf(personToGreet.FirstName, personToGreet.LastName)
	greeting := fmt.Sprintf("Hello %s %s", personToGreet.FirstName, personToGreet.LastName)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(greeting))

}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong!"))
}

func main() {

	// Create a new router
	rt := chi.NewRouter()

	// Endpoint
	rt.Get("/ping", pongHandler)
	rt.Post("/greetings", greetingsHandler)

	// Start server
	if err := http.ListenAndServe(":3000", rt); err != nil {
		panic(err)
	}

}
