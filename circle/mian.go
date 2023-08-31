package main

import (
	"gotest/circle/test_one"
	"gotest/circle/test_two"
)

//go:generate go run main.go
func main() {
	p1 := test_one.P1{}
	p1.SayP2()
	p2 := test_two.NewP2(p1)
	p2.P.SayP()
}
