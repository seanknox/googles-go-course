package main

import (
	"fmt"
)

func main() {
	allNumbers := []int{2, 3, 4, 5, 10, 11, 13, 15, 22}

	allSum := sum(allNumbers...)
	oddSum := oddNumberSum(sum, allNumbers...)
	fmt.Println(allSum)
	fmt.Println(oddSum)
}

func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

func oddNumberSum(sumFunc func(sumNumbers ...int) int, numbers ...int) int {
	var oddTotal []int
	for _, v := range numbers {
		if v%2 != 0 {
			oddTotal = append(oddTotal, v)
		}
	}

	return sumFunc(oddTotal...)
}
