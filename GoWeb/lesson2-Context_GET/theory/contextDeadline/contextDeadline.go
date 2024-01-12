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
	for i := 0; i < 15; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Operation in progress...")
		case <-ctx.Done():
			fmt.Println("Operation canceled")
		}
	}
}

func main() {

	deadline := time.Now().Add(5 * time.Second)
	ctx := context.Background()
	ctx, cancelFunc := context.WithDeadline(ctx, deadline)
	defer cancelFunc()

	go SimulateOperation(ctx)

	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println("End of the program")
}
