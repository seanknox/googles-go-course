package main

import (
	"fmt"
)

func main() {
	anonFuncChan()
	bufferChan()
}

func anonFuncChan() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}

func bufferChan() {
	c := make(chan int, 2)

	c <- 43
	c <- 44

	fmt.Println(<-c)
	fmt.Println(<-c)
}
