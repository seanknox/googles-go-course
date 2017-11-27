package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
	p1 := person{first: "Sarah", age: 44}
	p2 := person{first: "Tess", age: 24}
	p3 := person{first: "Sean", age: 37}
	p4 := person{first: "Michelle", age: 26}
	p5 := person{first: "Chris", age: 32}

	people := []person{p1, p2, p3, p4, p5}

	fmt.Println("Unsorted: ", people)

	sort.Sort(ByAge(people))
	fmt.Println("By age:", people)

	sort.Sort(ByName(people))
	fmt.Println("By name: ", people)

}
