package main

import "fmt"

func main() {
	// var 1firstName string // No puede arranacar con numeros
	// var lastName string // Correcta
	// var int age // pimero es el nombre y despues el tipo
	// 1lastName := 6 // No debe empezar con numero, debe ser tipo string si es un apellido
	// var driver_license = true // Esta bien
	// var person height int // No debe estar separado las palabras person_height
	// childsNumber := 2 // Podria tener sentido

	var fistName string    // Correcta
	var lastName string    // Correcta
	var age int            // Correcta
	lastName = "Smith"     // Correcta
	var driverLicense bool // Correcta
	var personHeight int   // Correcta
	var childsNumber int   // Correcta

	fmt.Print(fistName, lastName, age, lastName, driverLicense, personHeight, childsNumber)

}
