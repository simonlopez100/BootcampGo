package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {

		fmt.Println(err)
		return

	}
	f.Write([]byte("Hello world!!!\n"))

}
