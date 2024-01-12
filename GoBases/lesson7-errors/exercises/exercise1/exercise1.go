package main

import "fmt"

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return (*e).message
}

func main() {
	fmt.Println("Exercise1")

	var salary int = 1000

	if salary < 150000 {
		var err = CustomError{"Error: the salary entered does not reach the taxable minimum"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Must pay tax")
	}
}
