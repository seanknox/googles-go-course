package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "apple":
		fmt.Println("What a delicious fruit")
	case "banana":
		fmt.Println("Chalk full of potasium")
	default:
		fmt.Printf("Hmm, %v...not sure about that one.\n", os.Args[1])
	}
}
