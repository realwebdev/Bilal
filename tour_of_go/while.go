package main

import (
	"fmt"
)

func main() {

	// GO dont have while

	sum := 2

	for sum < 20 {

		sum += 1

		fmt.Println(sum)
	}

	fmt.Println(sum)
}
