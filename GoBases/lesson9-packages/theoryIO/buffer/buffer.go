package theoryIo

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	text := "lorem ipsum dolor sit amet"
	reader := strings.NewReader(text)

	// read by 5 bytes per time
	buf := make([]byte, 8)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		println("Data read:", string(buf[:n]))
	}

}
