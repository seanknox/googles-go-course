package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	bs, err := json.Marshal([]user{u1, u2, u3})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
