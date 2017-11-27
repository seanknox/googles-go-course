package main

import (
	"fmt"
)

type person struct {
	first_name       string
	last_name        string
	favorite_flavors []string
}

func main() {
	p1 := person{
		first_name:       "Sean",
		last_name:        "Knox",
		favorite_flavors: []string{"mint", "vanilla"},
	}
	p2 := person{
		first_name:       "Carcamo",
		last_name:        "Rox",
		favorite_flavors: []string{"yogurt"},
	}

	peeps := []person{p1, p2}

	for _, p := range peeps {
		fmt.Printf("%v %v's favorite ice cream flavors: ", p.first_name, p.last_name)
		for _, v := range p.favorite_flavors {
			fmt.Printf("%v\t", v)
		}
		fmt.Println()
	}

}
