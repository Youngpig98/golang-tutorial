package main

import (
	"fmt"
	"os"
)

func add(a, b int) int {
	return a + b
}

var g int = 100

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}

	a, b := 1, 2
	res := add(a, b)
	fmt.Println("a + b = ", res)
}
