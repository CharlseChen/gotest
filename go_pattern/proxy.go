package main

import "fmt"

type User interface {
	Echo()
}

type Man struct{}

func (m *Man) Echo() {
	fmt.Println("man")
}

type Woman struct{}

func (w *Woman) Echo() {
	fmt.Println("woman")
}

type UserProxy struct {
	User User
}

func (u *UserProxy) Echo() {
	u.User.Echo()
}

func NewUserProxy(user User) *UserProxy {
	return &UserProxy{
		User: user,
	}
}

func main() {
	proxy := NewUserProxy(new(Man))
	proxy.Echo()
}
