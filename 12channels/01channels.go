package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// channels are designed to synchronize data transmission between go routines

	ch := make(chan int) // by default, it is unbuffered channel

	wg.Add(2)
	// receiving go routine
	go func() {
		i := <-ch // data flowing out of the channel
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()

	// sending go routine
	go func() {
		ch <- 42 // data flowing into the channel
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}
