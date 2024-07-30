package exercise

import (
	"testing"
	"sync"
	"fmt"
	"golang.org/x/sync/singleflight"
)

func BenchmarkSyncOnce(b *testing.B) {

	a := 1
	b.RunParallel(func(pb *testing.PB) {
		var once sync.Once
		for pb.Next() {
			once.Do(func() {
				a++
			})
		}
	})
	fmt.Println(a)
}

func BenchmarkSingleFlight(b *testing.B) {
	var group singleflight.Group
	a := 1
	c := 1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			d, _, _ := group.Do("a", func() (interface{}, error) {
				a++
				return a, nil
			})
			c++
			fmt.Println(d, c)
		}
	})

}
