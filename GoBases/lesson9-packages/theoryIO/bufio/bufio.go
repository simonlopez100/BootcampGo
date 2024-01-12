package main

import (
	"bufio"
	"os"
)

func main() {

	rd := bufio.NewReader(os.Stdin)

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}

		println(line)
	}

}
