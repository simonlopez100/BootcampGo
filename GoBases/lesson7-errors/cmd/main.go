package main

import (
	"errors"
	"fmt"
	"lesson7-errors/internal"
	"lesson7-errors/internal/repository"
)

func main() {

	movieRepo := repository.NewMovieMap(3)

	movies := []internal.Movie{
		{
			Title:  "The Matrix",
			Year:   1998,
			Rating: 8.7,
		},
		{
			Title:  "The Lion King",
			Year:   1999,
			Rating: 8.7,
		},
		{
			Title:  "Pulp fiction",
			Year:   2000,
			Rating: 8.7,
		},
		{
			Title:  "Kill Bill",
			Year:   2004,
			Rating: 8.7,
		},
	}

	for _, mv := range movies {

		err := movieRepo.Save(&mv)
		err = errors.Unwrap(err)
		if err != nil {
			switch err {
			case repository.ErrLimitReached:
				fmt.Println("Code 01: limit reached")
			case repository.ErrInvalidMovie:
				fmt.Println("Code 02: Invalid movie")
			case repository.ErrConstraintsViolation:
				fmt.Println("Case 03: Constraint violation")
			default:
				fmt.Println("Case 04: Unknown error")
			}
		}
	}

	// mv := internal.Movie{
	// 	ID:     1,
	// 	Title:  "The Matrix",
	// 	Year:   1998,
	// 	Rating: 8.7,
	// }

	// err := movieRepo.Save(&mv)
	// if err != nil {
	// 	switch err {
	// 	case repository.ErrLimitReached:
	// 		fmt.Println("Code 01: limit reached")
	// 	case repository.ErrInvalidMovie:
	// 		fmt.Println("Code 02: Invalid movie")
	// 	case repository.ErrConstraintsViolation:
	// 		fmt.Println("Case 03: Constraint violation")
	// default:
	// 		fmt.Println("Case 04: Unkown error")
	// 	}
	// }

	// fmt.Printf("The movie is:%v", mv)
	// println()
	fmt.Printf("The movie repo is shown next: %v", movieRepo)

}
