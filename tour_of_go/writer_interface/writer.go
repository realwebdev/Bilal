package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Hello, Play")
	fmt.Fprintln(os.Stdout, "Hello, Play")
	io.WriteString(os.Stdout, "Hello, Play\n")
}
