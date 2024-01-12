package main

import "fmt"

func main() {
	const (
		age        int     = 23
		antique    int     = 2
		salary     float64 = 1200
		isEmployed bool    = true
	)

	bankLoadOk(age, antique, salary, isEmployed)
}

func bankLoadOk(age, antique int, salary float64, isEmployed bool) {

	if age > 22 && antique > 1 && isEmployed == true {
		if salary > 1000 {
			fmt.Println("Bank load with interest granted")
		} else {
			fmt.Println("Bank load without interest granted")
		}
	} else {
		fmt.Println("Bank load cant be granted")
	}

}
