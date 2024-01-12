package handlers

import (
	"encoding/json"
	"lesson3-POST/internal"
	"net/http"
)

func NewDefaultMovies()

type DefaultMovies struct {
	//Movies is a map of movies
	movies map[int]internal.Movie
	lastID int
}

type BodyReuestMovieJSON struct {
	Title     string  `json:"title"`
	Year      int     `json:"year"`
	Rating    float64 `json:"rating"`
	Published bool    `json:"published"`
}

func (d *DefaultMovies) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body BodyReuestMovieJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Bad request: invalid body"))
		}
	} // Add closing parenthesis and semicolon here
}
