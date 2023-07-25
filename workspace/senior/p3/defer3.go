// defer3.go
package main

import (
	"fmt"
	"os"
)

func test1() {
	fmt.Println("test")
}

func f() (result int) {
	defer func() {
		result *= 7
	}()

	return 6
}

func main() {
	fmt.Println("main start")
	defer test1()
	fmt.Println("main end")

	fmt.Println(f()) //f函数里的defer不会受到影响，因为在os.Exit被调用前它就被执行了
	os.Exit(0)
}
