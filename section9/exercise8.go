package main

import (
	"fmt"
)

// Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

// 		`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
// 		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
// 		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

// solution: https://play.golang.org/p/nTzSlRa9_A

func main() {
	favThings := map[string][]string{
		`knox_sean`:  []string{"bourbon", "bikes", "buddhism"},
		"bond_james": []string{`martinis`, `women`, `shaken not stirred`},
	}

	for k, v := range favThings {
		fmt.Printf("%v's favorite things:\n", k)
		for i, v := range v {
			fmt.Printf("%v: %v\n", i, v)

		}
	}

}
