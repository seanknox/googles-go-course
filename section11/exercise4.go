package main

import (
	"fmt"
)

func main() {
	x := struct {
		name    string
		age     int
		petType string
	}{
		name:    "Carcamo",
		age:     9,
		petType: "House Cat",
	}

	fmt.Println(x)

}
