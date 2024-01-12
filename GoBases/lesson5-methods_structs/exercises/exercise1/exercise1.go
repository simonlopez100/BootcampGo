package main

import "fmt"

// Define a struct called Product with the following fields: Name, Price, Description and Category.
type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Slice global variable called products of type Product.
var Products []Product

// Define method save to append a new product to the products slice.
func (p Product) save() {
	Products = append(Products, p)
}

// Define method getProducts that returns the products slice.
func (p Product) getProducts() []Product {
	fmt.Println("List of all the products:")
	return Products
}

// Define method getProduct that returns a product given its ID.
func (p Product) getProductById(id int) Product {
	fmt.Println("Product with ID", id, ":")
	return Products[id]
}

func main() {
	fmt.Println("Hello, exercise1!")

	product1 := Product{Id: 1, Name: "Product 1", Price: 100, Description: "Description 1", Category: "Category 1"}
	product2 := Product{Id: 2, Name: "Product 2", Price: 200, Description: "Description 2", Category: "Category 2"}
	product3 := Product{Id: 3, Name: "Product 3", Price: 300, Description: "Description 3", Category: "Category 3"}

	product1.save()
	product2.save()
	product3.save()

	fmt.Print(product1.getProducts())
	fmt.Print(product1.getProductById(2))
}
