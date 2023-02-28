package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// we will get a deadlock, if one of the parties is either sending or receiving and the other is unavailable
func main() {
	ch := make(chan int, 50) // buffered channel

	var doneCh = make(chan struct{}) // 0 memory allocation, signal only channel

	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch { // uses ,ok
			fmt.Println(i)
		}

		// same as above
		// for {
		// 	if i, ok := <- ch; ok {
		// 		fmt.Println(i)
		// 	} else {
		// 		break
		// 	}
		// }

		// for graceful shutdown of the go channel
		select {
		case entry := <-ch:
			fmt.Println(entry)
		case <-doneCh:
			break
		default:
			break
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 16
		close(ch) // prevents the deadlock, as we closed the channel
		wg.Done()
	}(ch)
	wg.Wait()

	doneCh <- struct{}{} // signalling for channel to shutdown
}
