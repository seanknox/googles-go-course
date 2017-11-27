package main

import (
	"fmt"
)

// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
func main() {

	numFunc := nestedFunc()

	fmt.Println(numFunc())
}

func nestedFunc() func() int {
	return func() int {
		return 99
	}
}
