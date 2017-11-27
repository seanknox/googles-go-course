package main

import (
	"fmt"
)

func main() {
	x := enclosed()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

	y := enclosed()
	fmt.Println(y())
	fmt.Println(y())
}

func enclosed() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
