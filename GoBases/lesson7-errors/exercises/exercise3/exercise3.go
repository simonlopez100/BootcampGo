package main

import (
	"errors"
	"fmt"
)

var errorSalaryLow error = errors.New("Error: salary is less than 10000")

func checkSalary(salary int) error {
	if salary <= 100000 {
		return errorSalaryLow
	}
	return nil
}

func main() {
	var salary int
	salary = 100000
	err := checkSalary(salary)
	if errors.Is(err, errorSalaryLow) {
		fmt.Println(err)
	}
}
