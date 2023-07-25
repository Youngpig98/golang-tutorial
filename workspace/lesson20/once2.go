package main

import (
	"fmt"
	"sync"
)

func print() {
	fmt.Println("test once")
}

func test(once *sync.Once) {
	once.Do(print)
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	size := 10
	wg.Add(size)

	/*启用size个goroutine，每个goroutine都调用once.Do(print)
	  最终print只会执行一次
	*/
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			test(&once)
		}()
	}
	/*等待所有goroutine执行完成*/
	wg.Wait()
	fmt.Println("end")
}
