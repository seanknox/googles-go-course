package main

import (
	"fmt"
)

func main() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan int)

	go send(evenChannel, oddChannel, quitChannel)

	receive(evenChannel, oddChannel, quitChannel)

	fmt.Println("about to quit!")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case e := <-e:
			fmt.Println("From the even channel: ", e)
		case o := <-o:
			fmt.Println("From the odd channel: ", o)
		case q := <-q:
			fmt.Println("From the quit channel: ", q)
			return
		}
	}
}
