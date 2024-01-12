package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	People   Person
}

func (e Employee) printEmployee() {
	fmt.Printf("Data for employee: %s, with id:  %d and birthday: %s \n", e.People.Name, e.People.ID, e.People.DateOfBirth)
	fmt.Printf("Employee ID: %d, Position: %s", e.ID, e.Position)
}

func main() {

	p1 := Person{ID: 1, Name: "Simon", DateOfBirth: "14/07/1992"}
	e1 := Employee{ID: 1, Position: "Software Developer", People: p1}
	fmt.Println(p1)
	e1.printEmployee()
}
