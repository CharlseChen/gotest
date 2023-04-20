package main

import "fmt"

type Tableware interface {
	UseTableware()
}
type chopstick struct {
}

func (k *chopstick) UseTableware() {
	fmt.Println("用筷子")
}

type spoon struct {
}

func (s *spoon) UseTableware() {
	fmt.Println("用汤勺")
}

type Me struct {
	tableware Tableware
}

func (t *Me) SetTableware(s Tableware) {
	t.tableware = s
}

func (t *Me) Fight() {
	t.tableware.UseTableware()
}

func main() {
	me := new(Me)
	me.SetTableware(new(chopstick))
	me.Fight()
}
