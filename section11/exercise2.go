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
		first_name: "Sean",
		last_name:  "Knox",
		favorite_flavors: []string{
			"mint",
			"vanilla",
		},
	}
	p2 := person{
		first_name:       "Carcamo",
		last_name:        "Rox",
		favorite_flavors: []string{"yogurt"},
	}

	peeps := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, v := range peeps {
		fmt.Printf("%v %v's favorite ice cream flavors: ", v.first_name, v.last_name)
		for _, v2 := range v.favorite_flavors {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}

}
