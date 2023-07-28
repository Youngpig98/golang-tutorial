// example1.go
package main

import "fmt"

type student struct {
	name string
}

// 最佳实践：尽量不要使用new，对于slice、map和channel，使用make创建
func main() {
	a := *new([]int)
	fmt.Printf("%T, %v\n", a, a == nil)

	b := *new(map[string]int)
	fmt.Printf("%T, %v\n", b, b == nil)

	c := *new(chan int)
	fmt.Printf("%T, %v\n", c, c == nil)

	d := &student{
		name: "wangyuyang",
	}
	fmt.Println("name is: ", d.name)

	mapPtr := new(map[string]int)
	fmt.Println("is map nil:", *mapPtr == nil)

	var iSlice []int32
	iSlice = append(iSlice, 10) //append函数可以对nil的slice进行扩容
	fmt.Printf("iSlice's cap is %d", cap(iSlice))
}
