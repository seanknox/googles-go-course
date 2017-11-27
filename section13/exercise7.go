package main

import (
	"fmt"
)

func main() {
	f := func() int {
		return 64
	}

	num := f()
	fmt.Println(num)
}
