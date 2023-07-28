package main

import (
	"fmt"
	"time"
)

// BackOffUntil2 goroutine 启停函数
func BackOffUntil2(stopCh chan struct{}, fn func()) {
	//经典 for select范式
	for {
		select {
		// 这里会读取到一个空struct
		case <-stopCh:
			return
		//在此处设置了一个计时器，意思每过一秒就做一些事情
		case <-time.After(1 * time.Second):
			fmt.Println("begin running fn()")
			//可以在这里调用一些函数，根据自己的需求对函数做更改
			fn()
		}
	}
}

// 定义一个公用的业务逻辑，在BackOffUntil中执行
func fn2() {
	fmt.Println("fn2 run")
}

func main() {
	stopCh := make(chan struct{}, 2) //定义一个缓冲区大小为2的channel，这样后面就不用每次都重新定义一个新的stopCh了

	//模拟3s后停止
	go func(chan struct{}) {
		select {
		case <-time.After(3 * time.Second):
			stopCh <- struct{}{} //传入一个空struct，而非前面的close channel
			fmt.Println("stopCh receives a empty struct")
		}
	}(stopCh)

	BackOffUntil2(stopCh, fn2)

	//模拟10s后停止
	go func(chan struct{}) {
		select {
		case <-time.After(10 * time.Second):
			stopCh <- struct{}{}
			fmt.Println("stopCh receives a empty struct")
		}
	}(stopCh)

	BackOffUntil2(stopCh, fn2)
}
