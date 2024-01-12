package main

import "fmt"

func main() {
	// Variable con el número del mes (puedes ajustar este valor según tus necesidades)
	mes := 7

	// Imprimir el nombre del mes
	imprimirNombreMes(mes)
}

func imprimirNombreMes(mes int) {
	// Utilizar un switch statement para asociar el número del mes con su nombre
	switch mes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("Número de mes no válido")
	}
}
