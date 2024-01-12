package main

import "fmt"

type Preferences struct {
	Foods  string
	Colors string
	Movies string
	Books  string
}

type Opeeration struct {
	Sum func(a, b int) int
}

type Person struct {
	Name  string
	Age   int
	Likes Preferences
}

type Circle struct {
	Radio float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radio * c.Radio
}

func main() {
	fmt.Println("Hello, Person struct!")

	person1 := Person{}
	person1.Name = "Simon"
	person1.Age = 31
	person1.Likes.Foods = "Pizza"

	fmt.Println(person1)

	c := Circle{Radio: 10}
	fmt.Println(c.Area())

}
