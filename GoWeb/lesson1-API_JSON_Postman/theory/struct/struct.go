package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func main() {

	bytes := []byte(`{"name": "John", "Age": 30, "Height": 2.0}`)

	var p Person
	if err := json.Unmarshal(bytes, &p); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p)
}
