package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name      string
	Price     int
	Published bool
}

var jsonData = `{"Name": "MacBook Air", "Price": 900, "Published": true}`

func main() {
	fmt.Println("Unmarshal")

	// convert string to []byte
	jsonDataByte := []byte(jsonData)

	// This example take only the values of the json data and save it to the variable p
	var p Product
	if err := json.Unmarshal(jsonDataByte, &p); err != nil {
		log.Fatal(err)
	}

	var mapJsonData map[string]interface{}
	if err := json.Unmarshal(jsonDataByte, &mapJsonData); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Take json data values and save it to variable p")
	fmt.Println(p)

	fmt.Println("Take json data complete and save it to variable mapJsonData")
	fmt.Println(mapJsonData)
}
