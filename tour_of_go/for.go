//for loops
package main

import (
	"fmt"
)

func foor() {
	//one looping construct in go has theree parts
	// init statement
	// condition expression
	// post statement// braces are always required but no parenthesis around components
	allsum := 0

	for i := 0; i < 10; i++ {

		allsum += i

		fmt.Println(allsum)
	}
	fmt.Printf("this is the total of all value %v\n", allsum)

}
func main() {

	foor()
}
