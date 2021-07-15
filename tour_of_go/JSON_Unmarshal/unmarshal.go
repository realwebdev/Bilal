package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Person struct {
		First string `json:"First"`
		Last  string `json:"Last"`
		Age   int    `json:"Age"`
	}
	s := `[{"First":"james","Last":"Bond","Age":32},{"First":"Miss","Last":"MonneyPenny","Age":32}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	people := []Person{}

	// var people []Person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("ALl of the data ", people)

	for i, v := range people {

		fmt.Println("\nPerson Number", i)
		fmt.Printf(v.First, v.Last, v.Age)
	}
}
