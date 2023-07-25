package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		fmt.Println("托尔斯泰")
		return nil
	})
	g.Go(func() error {
		fmt.Println("列宁")
		return nil
	})
	g.Go(func() error {
		fmt.Println("马克思")
		return nil
	})
	_ = g.Wait()
	fmt.Println("结束了")
}
