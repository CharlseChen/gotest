package main

import (
	_ "embed"
	"fmt"
)

//go:generate go run embed_example.go
//go:embed embed_example.go
//这里将embed_example.go中的内容写入到包一级的全局变量中
var src string

func main() {
	fmt.Println(src)
}
