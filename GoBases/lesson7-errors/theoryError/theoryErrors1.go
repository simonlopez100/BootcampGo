package main

import (
	"fmt"
	"strconv"

	"errors"
)

type CustomErr struct {
	Msg  string
	Code int
}

func (e *CustomErr) Error() string {
	return e.Msg + "" + strconv.Itoa(e.Code)
}

func main() {
	fmt.Println("Hello World")

	var err error = errors.New("Error message")
	fmt.Println("error:", err)

	var err11 error = &CustomErr{"Error message", 1}
	fmt.Println("error11:", err11)
}
