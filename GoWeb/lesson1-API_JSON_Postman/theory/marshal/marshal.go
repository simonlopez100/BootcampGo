package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name      string
	Price     int
	Publisher bool
}

func main() {
	fmt.Println("Marshal")

	p := Product{
		Name:      "MacBook Pro",
		Price:     1500,
		Publisher: true,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("from instance of struct to json data: %s\n", string(jsonData))
}
