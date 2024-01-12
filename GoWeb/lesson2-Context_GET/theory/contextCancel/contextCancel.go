package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func SimulateOperation(ctx context.Context) {
	fmt.Println("Simulate operation..")
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Operation in progress...")
		case <-ctx.Done():
			fmt.Println("Operation canceled")
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)

	go SimulateOperation(ctx)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancelFunc()
	fmt.Println("End of the program")
}
