package main

import (
	"fmt"
)

type burger int

var x burger
var y int

func main() {
	x := int(42)
	x = y
	fmt.Printf("Val: %v and Type: %T  ", x, x)
}
