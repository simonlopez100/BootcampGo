package main

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

func main() {
	jsonStream := `
	{"ProductID": "AXW23", "Price": 25.50}
	{"ProductID": "AXW24", "Price": 26.50}
	{"ProductID": "AXW25", "Price": 27.50}
	`

	myStreaming := strings.NewReader(jsonStream)
	myDecoder := json.NewDecoder(myStreaming)

	type myData struct {
		ProductID string
		Price     float64
	}

	for {
		var data myData

		if err := myDecoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		log.Println("DATA:", data.ProductID, data.Price)
	}
}
