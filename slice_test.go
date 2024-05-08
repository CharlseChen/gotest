package main

import (
	"testing"
	"fmt"
	"time"
)

func TestEquals(t *testing.T) {
	a1 := [2]int{1, 2}
	b1 := [2]int{1, 2}
	if a1 == b1 {
		fmt.Println("equal")
	}
	r := time.Date(0, 0, 0, 23, 55, 0, 0, time.Now().Location()).Unix()
	fmt.Printf("%d", r)
	ta := time.Unix(3600, 0)
	fmt.Printf("%s", ta.Format("2006-01-02 15-04-05"))
}
