package main

import "fmt"

// Methods and pointers example 1 -----------------------
type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) setFirstName(firstName string) {
	(*p).FirstName = firstName
}

//Methods and pointers example 2 -----------------------

type Location struct {
	City    string
	Country string
}

type Company struct {
	Name     string
	Location Location
}

func (c *Company) MigrateLocation(newLocation Location) {
	(*c).Location = newLocation
}

type Employee struct {
	Name     string
	Position string
	Company  Company
}

func (e Employee) Information() {
	fmt.Printf("%s works as %s at %s located in: %s, %s\n", e.Name, e.Position, e.Company.Name, e.Company.Location.City, e.Company.Location.Country)
}

func main() {
	fmt.Println("Pointers theory")

	// Se declara una variable number de tipo int con el valor 5
	var number int = 5

	// Se asigna a ptr la dirección de memoria de number
	var ptr *int = &number

	fmt.Println("memory address of ptr:", ptr)
	fmt.Println("value of ptr:", *ptr)
	fmt.Println()
	fmt.Println()

	// Methods and pointers example 1 -----------------------

	fmt.Println("Methods and pointers example 1")
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	person.setFirstName("Jane")

	fmt.Println("Complete name: ", person.FullName())
	fmt.Println()
	fmt.Println()

	// Methods and pointers example 2 -----------------------

	// Create employee
	e1 := Employee{
		Name:     "Simon Lopez",
		Position: "Software Developer",
		Company: Company{
			Name: "Mercado Libre",
			Location: Location{
				City:    "Medellin",
				Country: "Colombia",
			},
		},
	}

	// Print employee information
	e1.Information()

	newLocation := Location{
		City:    "Buenos Aires",
		Country: "Aregentina",
	}
	e1.Company.MigrateLocation(newLocation)
	e1.Information()

}
