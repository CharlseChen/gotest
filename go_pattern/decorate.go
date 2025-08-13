package main

import "fmt"

type Phone interface {
	Show()
}

type Realme struct {
}

func (r *Realme) Show() {
	fmt.Println("realme")
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {
	d.phone.Show()
}

type MoDecorator struct {
	Decorator
	Color func() string
}

func (m *MoDecorator) Show() {
	m.phone.Show()
	fmt.Println("贴膜")
}

func (m *MoDecorator) ShowColor() string {
	return func() string {
		return m.Color()
	}()
}
func NewMoDecorator(phone Phone) *MoDecorator {
	return &MoDecorator{Decorator: Decorator{phone: phone}, Color: func() string {
		return "red"
	}}
}

func main() {
	d := NewMoDecorator(new(Realme))
	d.Show()
	d.ShowColor()
}
