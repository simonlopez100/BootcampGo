package repository

import (
	"errors"
	"fmt"
	"lesson7-errors/internal"
)

var (
	ErrInvalidMovie         = errors.New("invalid movie")
	ErrConstraintsViolation = errors.New("Constrains validation ")
	ErrLimitReached         = errors.New("Storage for movies is full")
)

type MovieMap struct {
	movies map[int]internal.Movie
	lastID int
	limit  int
	count  int
}

func NewMovieMap(limit int) *MovieMap {
	return &MovieMap{
		movies: make(map[int]internal.Movie),
		limit:  limit,
	}
}

func (m *MovieMap) Save(movie *internal.Movie) (err error) {

	// Check limit
	if m.count >= m.limit {
		err := fmt.Errorf("%w: limit: %d", ErrLimitReached, m.limit)
		return err
	}

	// Check if movie name exist, cant save movie without name
	if movie.Title == "" {
		err := ErrInvalidMovie
		return err
	}

	// Chek if movie has a rate between 0 and 10
	if !(movie.Rating >= 0 && movie.Rating <= 10) {
		err := ErrInvalidMovie
		return err
	}

	// Check if movie was created between 1990 and 2100
	if movie.Year < 1900 || movie.Year > 2100 {
		err := ErrInvalidMovie
		return err
	}

	// Check if the movie already exist in the storage
	for _, m := range m.movies {
		if m.Title == movie.Title {
			err := ErrConstraintsViolation
			return err
		}
	}

	// Add automatic id when a movie es saved in the storage looking for the last id
	// increment id
	(*m).lastID++
	// set id
	movie.ID = (*m).lastID
	// save
	m.movies[movie.ID] = *movie

	// increment count
	m.count++

	return

}
