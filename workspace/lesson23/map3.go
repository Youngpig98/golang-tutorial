package main

import (
	"fmt"
	"sync"
)

var m = map[int]int{} //初始化
var mu = sync.Mutex{}

func changeMap(key int) {
	//需要加锁
	mu.Lock()
	defer mu.Unlock()
	m[key] = 1
}

func main() {
	/*统计字符串里每个字符出现的次数*/
	var wg sync.WaitGroup
	size := 2
	wg.Add(size)

	for i := 0; i < size; i++ {
		i := i
		go func() {
			defer wg.Done()
			changeMap(i)
		}()
	}
	wg.Wait()

	for key, value := range m {
		fmt.Printf("%d: %d\n", key, value)
	}
}
