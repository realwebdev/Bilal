package main

import (
	"fmt"
)

func pow(x, n, lim float64) float64 {
	// these are the kind of syntax tricky for understanding understand
	if v := x * n; v > lim {
		return v

	}
	return lim
}

func main() {

	fmt.Println(
		pow(3, 9, 10),
		pow(3, 2, 20),
	)
}
