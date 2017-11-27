package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go func() {
		c <- "Hello world!"
	}()

	fmt.Println(<-c)
}
