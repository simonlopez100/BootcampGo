package main

import "fmt"

func calculateSalary(workedHours float64, hourlyRate float64) (float64, error) {
	if workedHours <= 80 {
		return 0, fmt.Errorf("Error: the worker cannot have worked less than 80 hours per month")
	} else if hourlyRate <= 0 {
		return 0, fmt.Errorf("Error: hourly rate cant be less than or equal to 0")
	}
	return workedHours * hourlyRate, nil
}

func taxes(salary float64) (float64, float64, error) {
	if salary <= 150000 {
		return 0, 0, fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %f", salary)
	}

	taxes := 0.10
	salary = salary - (salary * taxes)
	return salary, taxes, nil
}

func main() {
	fmt.Println("Exercise5")

	var workedHours float64 = 180
	var hourlyRate float64 = 1000

	Salary, Err := calculateSalary(workedHours, hourlyRate)
	if Err != nil {
		fmt.Println(Err)
	} else {
		fmt.Println("Your base salary is:", Salary)
	}

	salary, taxes, Err := taxes(Salary)
	if Err != nil {
		fmt.Println(Err)
	} else {
		fmt.Printf("Salary: %f\nTaxes: %f\n", salary, taxes)
	}
}
