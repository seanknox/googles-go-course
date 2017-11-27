package main

import (
	"fmt"
)

const (
	current       = 2017 + iota
	last          = 2017 - iota
	twoyearsago   = 2017 - iota
	threeyearsago = 2017 - iota
)

func main() {
	fmt.Println(current, last, twoyearsago, threeyearsago)
}
