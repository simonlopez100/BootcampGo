package main

import (
	"encoding/json"
	"os"
)

func main() {
	myEncoder := json.NewEncoder(os.Stdout)

	type MyData struct {
		ProductID string
		Price     float64
	}

	data := MyData{
		ProductID: "XASD",
		Price:     25.50,
	}

	myEncoder.Encode(data)
}
