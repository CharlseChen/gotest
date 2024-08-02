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

//go:generate go run strategy.go
func main() {
	me := &Me{
		tableware: Tableware(&chopstick{}),
	}
	me.Fight()
}
