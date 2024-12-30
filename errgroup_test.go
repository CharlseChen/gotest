package main

import (
	"testing"
	"golang.org/x/sync/errgroup"
	"time"
	"errors"
)

/*
在errgroup中，其中一个协程报错，不会立即结束等待状态，会等所有协程都运行完
*/
func TestErrgroup(t *testing.T) {
	var g errgroup.Group
	g.Go(func() error {
		time.Sleep(time.Second)
		return errors.New("err1223")
	})
	g.Go(func() error {
		i := 0
		for i < 20 {
			i++
			t.Log(i)
			time.Sleep(time.Second)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		t.Fatal(err.Error())
	}
}
