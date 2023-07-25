package main

import "fmt"

func main() {
	slice := make([]int, 3, 10) //大小为3，容量为10
	/*下标访问切片*/
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%d\n", i, slice[i])
	}

	/*range迭代访问切片*/
	for index, value := range slice {
		fmt.Printf("slice[%d]=%d\n", index, value)
		if index == 0 {
			value = 100
		}
	}

	//发现在range中并不会实际修改切片的值
	fmt.Println(slice[0])
}
