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

//建一个手机相关的装饰器，并包含手机的功能
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {
	d.phone.Show()
}

// 基于现有的装饰器添加新功能
type MoDecorator struct {
	Decorator
	Color func() string
}

// 对现有功能进行扩展
func (m *MoDecorator) Show() {
	m.phone.Show()
	fmt.Println("贴膜")
}

// 新增新的功能
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
