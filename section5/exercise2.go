package main

import (
	"fmt"
)

func main() {

	equal := (4 == 4)
	lteq := (4 <= 5)
	gteq := (8 <= 8)
	noteq := ("hotdog" != "water")
	lt := (1 < 2)
	gt := (2 > 5)

	fmt.Println(equal, lteq, gteq, noteq, lt, gt)
}
