package exercise

import (
	"testing"
	"sync"
	"fmt"
	"golang.org/x/sync/singleflight"
)

func BenchmarkSyncOnce(b *testing.B) {
	var once sync.Once
	a := 1
	once.Do(func() {
		a++
	})
	fmt.Println(a)
}

func BenchmarkSingleFlight(b *testing.B) {
	var group singleflight.Group
	a := 1
	c, _, shared := group.Do("a", func() (interface{}, error) {
		a++
		return a, nil
	})
	fmt.Println(c, shared)
}
