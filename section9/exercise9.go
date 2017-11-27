package main

import (
	"fmt"
)

// Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
// solution: https://play.golang.org/p/_CkxAhRrDJ

// Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
// solution: https://play.golang.org/p/TYl5EbjoeC

func main() {
	m := map[string][]string{
		`knox_sean`:  []string{"bourbon", "bikes", "buddhism"},
		"bond_james": []string{`martinis`, `women`, `shaken not stirred`},
	}

	m[`jones_rad`] = []string{"foo", "bar", "baz"}
	delete(m, `bond_james`)

	for k, v := range m {
		fmt.Printf("%v's favorite things:\n", k)
		for i, v2 := range v {
			fmt.Printf("%v: %v\n", i, v2)

		}
	}

}
