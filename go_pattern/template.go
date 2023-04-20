package main

import "fmt"

type RushCar interface {
	Water()
	Bubble()
	Rush()
	Wipe()
}

type template struct {
	b RushCar
}

func (t *template) DoRushCar() {
	t.b.Water()
	t.b.Bubble()
	t.b.Rush()
	t.b.Wipe()
}

type BusRush struct {
	template
}

func (b *BusRush) Water() {
	fmt.Println("卡车冲水")
}

func (b *BusRush) Bubble() {
	fmt.Println("卡车打泡沫")
}

func (b *BusRush) Rush() {
	fmt.Println("卡车冲洗")
}

func (b *BusRush) Wipe() {
	fmt.Println("卡车擦干水")
}

func NewBusRush() *BusRush {
	B := &BusRush{}
	B.b = B
	return B
}

func main() {
	busRush := NewBusRush()
	busRush.DoRushCar()
}
