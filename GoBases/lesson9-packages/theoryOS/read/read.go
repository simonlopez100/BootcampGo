package main

import (
	"fmt"
	"os"
)

func main() {

	// open file

	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	buf := make([]byte, 8)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
