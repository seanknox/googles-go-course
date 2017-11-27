package main

import (
	"fmt"
)

type person struct {
}

func (p *person) speak() {
	fmt.Println("I'm saying somethin' over here")
}

type boozehound struct {
}

func (b *boozehound) speak() {
	fmt.Println("Imma let you finish but first I, I, I, uh, I gotta, um...")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{}
	saySomething(&p)

	b := boozehound{}
	saySomething(&b)
}
