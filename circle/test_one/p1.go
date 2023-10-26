package test_one

import (
	"gotest/circle/test_two"
	"fmt"
)

type P1 struct{}

func NewP1() *P1 {
	return &P1{}
}

func (p P1) SayP() {
	fmt.Println("This is P1")
}

func (p P1) SayP2() {
	t := test_two.NewP2(p)
	t.SayP2()
}
