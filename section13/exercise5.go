package main

import (
	"fmt"
	"math"
)

// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type square struct {
	x float64
	y float64
}

type shape interface {
	area() float64
}

type circle struct {
	r float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func (s square) area() float64 {
	return (s.x * s.y)
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r

}

func main() {
	s1 := square{
		x: 10,
		y: 10,
	}

	c1 := circle{
		r: 10.2,
	}

	info(s1)
	info(c1)
}
