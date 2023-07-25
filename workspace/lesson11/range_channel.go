package main

import (
	"fmt"
	"time"
)

func addData(ch chan int) {
	/*
		每1秒往通道里发送一次数据
	*/
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	//数据发送完毕
	close(ch)
}

func main() {
	ch := make(chan int, 10)

	go addData(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
