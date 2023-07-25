package main

import "fmt"

func main() {
	a := 10
	b := 20

	defer func(a int, b int) {
		fmt.Printf("a=%d,b=%d", a, b)
	}(a, b)

	a = 111
	b = 222
	fmt.Printf("a=%d,b=%d", a, b)
}
