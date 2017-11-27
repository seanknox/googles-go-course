package main

import (
	"fmt"
)

// Using a COMPOSITE LITERAL:
// create a SLICE of TYPE int
// assign 10 VALUES
// Range over the slice and print the values out.
// Using format printing
// print out the TYPE of the slice
// solution: https://play.golang.org/p/sAQeFB7DIs

func main() {
	x := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 101}

	for i, v := range x {
		fmt.Printf("%v: %v\n", i, v)
	}

	fmt.Printf("%T\n", x)

}
