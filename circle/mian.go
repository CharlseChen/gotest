package main

import "fmt"

//func main() {
//	p1 := test_one.P1{}
//	p1.SayP2()
//	p2 := test_two.NewP2(p1)
//	p2.P.SayP()
//}

type m[T int | string | float64] map[T]interface{}

//go:generate go run main.go
func main() {
	var ts = make(m[string], 10)
	ts["test"] = 1
	ts["test2"] = 2
	fmt.Println(ts)
}
