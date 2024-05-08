package main

import (
	_ "embed"
	"fmt"
)

//go:generate go run embed_example.go
//go:embed embed_example.go
var src string

func main() {
	fmt.Println(src)
}
