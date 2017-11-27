package main

import (
	"fmt"
)

// Hands on exercise
// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x, y, z)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 84, "doubled"
}
