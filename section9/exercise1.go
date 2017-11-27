package main

import (
	"fmt"
)

// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array
// solution: https://play.golang.org/p/tD0SzV3hdf

func main() {
	meValues := [5]int{11, 22, 33, 99, 101}
	for i, v := range meValues {
		fmt.Printf("%v: %v\n", i, v)
	}
	fmt.Printf("%T\n", meValues)
}
