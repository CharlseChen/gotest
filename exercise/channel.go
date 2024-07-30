package exercise

import (
	"fmt"
	"time"
)

//go:generate go run channel.go
func ChannelTest() {
	c := make(chan int)
	go writeChan(c, 888)
	time.Sleep(1 * time.Second)
	if _, ok := <-c; ok {
		fmt.Println(ok)
	} else {
		fmt.Println("channel closed")
	}

}

func writeChan(c chan int, x int) {
	fmt.Println("1111", x)
	c <- x
	fmt.Println(22222)
	close(c)
	fmt.Println(x)
}
