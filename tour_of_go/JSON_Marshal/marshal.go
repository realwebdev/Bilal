package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "james",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "MonneyPenny",
		Age:   32,
	}

	people := []person{
		p1,
		p2,
	}
	fmt.Println(people)

	//marshal take type of any type and return a slice of byte and error

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(string(bs))

}
