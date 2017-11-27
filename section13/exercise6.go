package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hi there")
	}()
}
