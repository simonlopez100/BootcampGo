package main

import "fmt"

func main() {
	fmt.Println("Hello, worked!")
	var salary, hoursWorked float64 = payment(3200, "A")
	fmt.Printf("Employee worked: %.2f hours and get payment for %.2f dollars\n", hoursWorked, salary)

}

func payment(minutesWorked float64, category string) (salary, hoursWorked float64) {
	var salaryPerHour float64
	switch category {
	case "A":
		salaryPerHour = 3000
	case "B":
		salaryPerHour = 1500
	case "C":
		salaryPerHour = 1000
	default:
		salaryPerHour = 0
	}
	hoursWorked = float64(minutesWorked) / 60
	salary = salaryPerHour * hoursWorked

	return salary, hoursWorked
}
