package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	numberOfEmployees := 0
	for _, employe := range employees {
		if employe > 21 {
			numberOfEmployees++
		}
	}
	fmt.Printf("Número total de empleados mayores de 21 años: %d\n", numberOfEmployees)

	employees["Federico"] = 21
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}
