package main

import (
	"context"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "my-key", "My special value")
	ctx = context.WithValue(ctx, "my-key-int", 5)
	user := User{
		Name: "Simon",
		Age:  31,
	}
	ctx = context.WithValue(ctx, "user", user)

	viewContext(ctx)
}

func viewContext(ctx context.Context) {
	fmt.Printf("My value is %s\n", ctx.Value("my-key"))
	fmt.Printf("My value is %v\n", ctx.Value("my-key-int"))
	fmt.Printf("My value is %v\n", ctx.Value("user"))
}
