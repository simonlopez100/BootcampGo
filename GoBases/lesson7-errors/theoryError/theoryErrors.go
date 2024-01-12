package main

import (
	"errors"
	"fmt"
)

// CustomError is a custom error struct to use method Error()
type CustomError struct {
	message string
	code    int
}

// Implement the Error() method from the CustomError struct
func (e *CustomError) Error() string {
	return (*e).message
}

//handle erros

// --------Handle errors---------
var ErrDivideByZero = errors.New("cannot divide by zero")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main1() {

	// Using CustomError struct and method Error()
	var err1 = CustomError{"Error message", 1}
	fmt.Println("error1:", err1.Error())

	// Using internal errors package
	err2 := errors.New("Error message")
	fmt.Println("error2:", err2)

	// Using internal fmt package
	err3 := fmt.Errorf("Error message")
	fmt.Println("error3:", err3)

	// diferent to errors.New, fmt.Errorf allow us to add a previus error
	// that let us create a chain of errors
	errBase := errors.New("Error base")
	errWrap := fmt.Errorf("Error wrap: %w", errBase)
	fmt.Println("errorWrap:", errWrap)

	// Is() from package errors let us check if an error is of a specific type y the wrap tree and return true or false
	fmt.Println(errors.Is(errWrap, errBase))

	// As() from package errors let us check if an error matches the target and return previus error
	errBase2 := &CustomError{"Error base", 1}
	errWrap2 := fmt.Errorf("Error wrap: %w", errBase2)

	var errAs *CustomError
	if errors.As(errWrap2, &errAs) {
		fmt.Println("errorAs:", errAs)
	}

	// --------Handle errors---------
	result, err4 := divide(10, 0)
	if err4 != nil {
		fmt.Println(err4)
		return
	}
	fmt.Println("result:", result)

}
