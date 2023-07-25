package main

import "fmt"

func main() {
	a := []int{1, 2}
	fmt.Println("length of a: ", len(a))
	fmt.Println("cap of a: ", cap(a))
	b := append(a, 3)
	fmt.Println("length of b: ", len(b))
	fmt.Println("cap of b: ", cap(b))

	c := append(b, 4)
	fmt.Println("length of b: ", len(b))
	fmt.Println("cap of b: ", cap(b))
	fmt.Println("length of c: ", len(c))
	fmt.Println("cap of c: ", cap(c))

	d := append(b, 5)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a, b, c[3], d[3])

	//for i := 0; i < 6; i++ {
	//	go func() {
	//		print(i)
	//	}()
	//}
	//
	//time.Sleep(1)
}
