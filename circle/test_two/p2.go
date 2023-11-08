package test_two

import "fmt"

type P interface {
	SayP()
}

type P2 struct {
	P P
}

func NewP2(p P) *P2 {
	return &P2{
		P: p,
	}
}

func (p *P2) SayP() {
	p.P.SayP()
}

func (p *P2) SayP2() {
	fmt.Println("This is P2")
}
