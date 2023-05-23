package main

import "fmt"

type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {
	d.phone.Show()
}

type Huawei struct {
}

func (h *Huawei) Show() {
	fmt.Println("huawei")
}

type Realme struct {
}

func (r *Realme) Show() {
	fmt.Println("realme")
}

type MoDecorator struct {
	Decorator
}

func (m *MoDecorator) Show() {
	m.phone.Show()
	fmt.Println("贴膜")
}

func NewMoDecorator(phone Phone) *MoDecorator {
	return &MoDecorator{Decorator{phone: phone}}
}

func main() {
	d := NewMoDecorator(new(Realme))
	d.Show()
}
