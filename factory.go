package main

import "fmt"

type fruit interface {
	Echo()
}

type Factory interface {
	CreateFruit() fruit
}

type AppleFactory struct {
}

func (a AppleFactory) CreateFruit() fruit {
	return fruit(&Apple{})
}

type Apple struct {
}

func (l *Apple) Echo() {
	fmt.Println("apple")
}

func main() {
	f := &AppleFactory{}
	f.CreateFruit().Echo()
}
