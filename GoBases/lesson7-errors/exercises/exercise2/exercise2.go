package main

import (
	"errors"
	"fmt"
)

type salaryError struct {
	message string
}

func (e *salaryError) Error() string {
	return e.message
}

func main() {
	fmt.Println("Exercise 2")

	// salary := 8

	err3 := &salaryError{message: "Error: salary is less than 2000"}
	err4 := &salaryError{message: "Error: salary is less than 10000"}

	if errors.Is(err3, err4) {
		fmt.Println("cosa")
	}

}
