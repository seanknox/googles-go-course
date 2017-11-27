package main

import (
	"fmt"
)

func main() {
	james := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	combined := [][]string{james, mp}

	for i, v := range combined {
		fmt.Printf("%v: %v", i, v)
	}
	fmt.Printf("%T", combined)
}
