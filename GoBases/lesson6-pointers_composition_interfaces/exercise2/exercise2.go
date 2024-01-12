package main

import "fmt"

const (
	Small  = "Small"
	Medium = "Medium"
	Large  = "Large"
)

type Product interface {
	Price() float64
}

type SmallSize struct {
	Cost float64
}

func (s SmallSize) Price() float64 {
	return s.Cost
}

type MediumSize struct {
	Cost float64
}

func (m MediumSize) Price() float64 {
	return m.Cost + m.Cost*0.03
}

type LargeSize struct {
	Cost float64
}

func (l LargeSize) Price() float64 {
	return l.Cost + l.Cost*0.06 + 2500
}

func ProductFactory(productType string, cost float64) Product {
	switch productType {
	case Small:
		return SmallSize{cost}
	case Medium:
		return MediumSize{cost}
	case Large:
		return LargeSize{cost}
	default:
		return nil
	}
}

func main() {
	fmt.Println("Exercise 2")

	ant := SmallSize{Cost: 1000}
	fmt.Println("The ant price is: ", "$", ant.Price())

	dog := MediumSize{Cost: 1500}
	fmt.Println("The dog price is: ", "$", dog.Price())

	elephant := LargeSize{Cost: 2500}
	fmt.Println("The elephant price is: ", "$", elephant.Price())

}
