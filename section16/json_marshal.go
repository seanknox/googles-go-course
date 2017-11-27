package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Tess","Last":"R","Age":25},{"First":"Sarah","Last":"Mei","Age":44}]`
	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Printf("Person %v:\n", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
