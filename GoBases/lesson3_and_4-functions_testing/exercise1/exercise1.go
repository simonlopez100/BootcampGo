package exercise1

import "fmt"

func main() {

	var salary1 float64 = 151000
	var salary2 float64 = 51000

	netSalary1, taxes1 := CalculateSalary(salary1)
	fmt.Printf("Net salary: %.2f, Taxes: %.2f%%\n", netSalary1, taxes1)
	netSalary2, taxes2 := CalculateSalary(salary2)
	fmt.Printf("Net salary: %.2f, Taxes: %.2f%%\n", netSalary2, taxes2)
}

func CalculateSalary(salary float64) (netSalary float64, taxes float64) {

	switch {
	case salary > 150000:
		taxes = 27
	case salary > 50000:
		taxes = 17
	default:
		taxes = 0
	}

	netSalary = salary - (salary * taxes / 100)
	return netSalary, taxes
}
