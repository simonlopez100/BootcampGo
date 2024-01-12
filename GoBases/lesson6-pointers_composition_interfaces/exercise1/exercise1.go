package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name        string
	lastName    string
	DNI         int
	DateOfStart time.Time
}

func (s Student) Details() {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Last Name: ", s.lastName)
	fmt.Println("DNI: ", s.DNI)
	fmt.Println("Date of Start: ", s.DateOfStart)
}

func main() {
	fmt.Println("Exercise 1")
	fmt.Println()

	student := Student{
		Name:        "John",
		lastName:    "Doe",
		DNI:         12345678,
		DateOfStart: time.Now(),
	}

	student1 := Student{
		Name:        "Jane",
		lastName:    "Cameron",
		DNI:         87654321,
		DateOfStart: time.Now(),
	}
	student.Details()
	fmt.Println("-------------------------------")
	student1.Details()
}
