package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func player(name string, ch chan int) {
	defer wg.Done()

	for {
		ball, ok := <-ch // fetch the value from channel

		if !ok {
			fmt.Printf("channel is closed! %s wins!\n", name)
			return
		}

		n := rand.Intn(100)

		if n%10 == 0 {
			close(ch)
			fmt.Printf("%s misses the ball! %s loses!\n", name, name)
			return
		}

		ball++
		fmt.Printf("%s receives ball %d\n", name, ball)
		ch <- ball

	}
}

var wg sync.WaitGroup

func main() {

	wg.Add(2) //  2 jobs to finish

	//go func(id string) {
	//	fmt.Println(id)
	//	wg.Done()
	//}("hello")
	//go say("world")

	ch := make(chan int, 0) // unbuffered channel

	//ch <- 0 //This will cause the deadlock in main, because nobody will fetch the value from the channel

	go player("wang", ch)

	go player("chong", ch)

	ch <- 0

	wg.Wait() // wait for all jobs to be finished

}
