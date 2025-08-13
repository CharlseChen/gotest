package main

import "fmt"

// 旧的电池
type V5 interface {
	Use5V()
}

type ApplePhone struct {
	v V5
}

func NewPhone(v V5) *ApplePhone {
	return &ApplePhone{v}
}

func (p *ApplePhone) Charge() {
	fmt.Println("Phone进行充电...")
	p.v.Use5V()
}

//=======================================================================================================================
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")
	a.v220.Use220V()
}

func main() {
	iphone := NewPhone(&Adapter{new(V220)})
	iphone.Charge()
}
