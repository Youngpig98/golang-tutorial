package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4, 5}

	var dst2 []*int
	//这里i是一个临时变量，因此append函数添加的其实都是临时变量i的地址
	//而不是src数组中每个变量的地址
	for _, i := range src {
		dst2 = append(dst2, &i)
	}

	//因此这里会输出55555
	for _, p := range dst2 {
		fmt.Print(*p)
	}

}
