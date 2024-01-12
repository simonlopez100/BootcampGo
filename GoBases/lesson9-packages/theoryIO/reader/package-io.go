package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	text := "lorem ipsum dolor sit amet"
	reader := strings.NewReader(text)

	buf := make([]byte, 30)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	println("Data read:", string(buf[:n]))

	reader.Seek(0, io.SeekStart)

	buf = make([]byte, 30)
	n, err = reader.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	println("Data read:", string(buf[:n]))

}
