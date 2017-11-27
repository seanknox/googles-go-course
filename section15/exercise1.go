package main

import "fmt"

// Create a value and assign it to a variable.
// Print the address of that value.

func main() {
	x := "a wonderful string"
	y := "a wonderful stringz"
	z1 := 99
	z2 := 99
	z3 := 99
	fmt.Println(&x, &y, &z1, &z2, &z3)
}
