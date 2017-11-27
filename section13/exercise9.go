package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 2, 3, 2, 2}
	fmt.Printf("Sum of all numbers: %v\n", sum(nums...))
	fmt.Printf("Sum of even numbers: %v\n", evenSum(sum, nums...))
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func evenSum(f func(otherFuncNums ...int) int, nums ...int) int {
	var evenTotal []int
	for _, v := range nums {
		if v%2 == 0 {
			evenTotal = append(evenTotal, v)
		}
	}

	return f(evenTotal...)
}
