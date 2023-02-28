package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int) { // receive-only "from" channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch) // casting from bidirectional channel to uni-directional channel

	go func(ch chan<- int) { // send-only "to" channel
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}
