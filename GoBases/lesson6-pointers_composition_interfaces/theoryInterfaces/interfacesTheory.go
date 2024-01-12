package main

import "fmt"

// Shape interface to multiple geometric shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// SHAPE 1 - CIRCLE -----------------------
// Circle struct to represent a circle
type Circle struct {
	Radius float64
}

// Area method to calculate the area of a circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Perimeter method to calculate the perimeter of a circle
func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// SHAPE 2 - RECTANGLE -----------------------
// Rectangle struct to represent a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method to calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method to calculate the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ShowDetails function that use the Shape interface to show the details of a shape
func ShowDetails(s Shape) {
	fmt.Println("Shape details:", s)
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main() {
	fmt.Print("Interfaces theory\n")

	// Creation of instances of Circle and Rectangle
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}

	// Using function from the Shape interface
	ShowDetails(&c)
	ShowDetails(r)
	fmt.Println()

	fmt.Println("Type Assertion")

	// Type Assertion
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

}
