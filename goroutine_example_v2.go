package main

import (
	"fmt"
	"context"
)

//go:generate go run goroutine_example_v2.go
func main() {
	ctx := context.Background()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err", err)
		}
	}()
	go func(ctx context.Context) {
		panic("错误")
	}(ctx)
	for {
	}
}
