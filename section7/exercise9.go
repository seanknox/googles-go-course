package main

import (
	"fmt"
	"os"
)

func main() {

	switch os.Args[1] {
	case "baseball":
		fmt.Println("America's old past time!")
	case "soccer":
		fmt.Println("The world's game")
	case "coffee":
		fmt.Println("The richest and arguably best sport")
	default:
		fmt.Println("Probably a good sport, but I'm not familiar with it!")
	}
}
