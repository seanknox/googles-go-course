package main

import (
	"fmt"
)

// create a person struct
// create a func called “changeMe” with a *person as a parameter
// change the value stored at the *person address
// important
// to dereference a struct, use (*value).field
// p1.first
// (*p1).first
// “As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
// https://golang.org/ref/spec#Selectors

// create a value of type person
// print out the value
// in func main
// call “changeMe”
// in func main
// print out the value

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{
		first: "Wayne",
	}
	fmt.Println(p1.first)
	changeMe(&p1)
	fmt.Println(p1.first)
}

func changeMe(p *person) {
	(*p).first = "Bruce"
}
