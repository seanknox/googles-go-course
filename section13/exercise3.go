package main

import (
	"fmt"
)

func main() {
	defer lateFunc()
	earlyFunc()
}

func lateFunc() {
	fmt.Println("...this happens last.")
}

func earlyFunc() {
	fmt.Println("This comes first.")
}
