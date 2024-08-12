package main

import (
	"testing"
	"golang.org/x/sync/errgroup"
	"time"
	"errors"
)

func TestErrgroup(t *testing.T) {
	var g errgroup.Group
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		return nil
	})
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		return nil
	})
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		return errors.New("err1223")
	})
	if err := g.Wait(); err != nil {
		t.Fatal(err.Error())
	}
}
