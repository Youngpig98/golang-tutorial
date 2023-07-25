package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	member int
}

var instance *Singleton

var mu sync.Mutex

var once sync.Once

func getInstance() *Singleton {
	/*
	   通过sync.Once实现单例，只会生成一个Singleton实例
	*/
	once.Do(func() { //闭包
		fmt.Println("once")
		instance = &Singleton{}
		instance.member = 100
	})
	return instance
}

func getInstance2() *Singleton {
	if instance == nil {
		mu.Lock()
		if instance == nil {
			instance = &Singleton{
				member: 15,
			}
		}
		mu.Unlock()
	}
	return instance
}

func main() {
	var wg sync.WaitGroup
	size := 10
	wg.Add(size)
	/*
	   多个goroutine同时去获取Singelton实例
	*/
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			instance = getInstance()
		}()
	}
	wg.Wait()
	fmt.Println(instance.member)
	fmt.Println("end")
}
