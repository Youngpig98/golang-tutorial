package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32
var mtx sync.Mutex // method one
var ch2 = make(chan int, 1)
var wg sync.WaitGroup

func MutexIncCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		mtx.Lock()
		counter++
		mtx.Unlock()
	}
}

func UnsafeIncCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		counter++
	}
}

func AtomicIncCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func ChannelIncCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		count := <-ch2
		count++
		ch2 <- count
	}

}

func raceCondition() {

	//race condition
	//go UnsafeIncCounter()
	//go UnsafeIncCounter()

	//go MutexIncCounter()
	//go MutexIncCounter()

	//go AtomicIncCounter()
	//go AtomicIncCounter()

	ch2 <- 0

	go ChannelIncCounter()
	go AtomicIncCounter()
}

func main() {
	wg.Add(2) //  2 jobs to finish

	//go func(id string) {
	//	fmt.Println(id)
	//	wg.Done()
	//}("hello")
	//go say("world")

	raceCondition()

	wg.Wait() // wait for all jobs to be finished

	fmt.Println("counter is ", counter)
	fmt.Println(<-ch2)
}
