package main

import "fmt"

func main() {
	var lastName string = "Smith" // Esta bien
	var age int = 35              // Se debe pasar como numero (sin comillas)
	boolean := false              // "Se debe pasar sin comillas"
	var salary float64 = 45857.90 // debe ser de tipo float64
	var firstName string = "Mary" // Esta bien

	fmt.Print(lastName, age, boolean, salary, firstName)
}
