package main

import (
	"fmt"
)

type sean int

var x sean
var y = "James Bond"
var z = true

func main() {
	fmt.Printf("%v %T", x, x)
	x = 42
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
