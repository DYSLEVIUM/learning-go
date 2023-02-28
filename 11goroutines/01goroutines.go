package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // designed to syncronize multiple go routines

// go run -race 11goroutines/01goroutines.go to detect race conditions
func main() {
	/*
		green threads are different from os threads. OS threads have their own individual function call stack dedicated to the execution of that thread, traditionally they are very large; about 1MB of RAM, take a lot of time for application to set up. This brings the concept of thread-pooling as creating and destruction of threads is very expensive in most programming language

		So we make a abstraction of a thread called go-routines. Inside the go runtime, we got a scheduler that maps these go-routines onto these OS threads for a period of time. And the scheduler will take turns with every CPU threads that is available and assign the different go-routines a certain amount of processing time on those threads, and we don't have to interact with the low level OS threads.

		So the advantage is that go-routines can start with very very small stack sizes, as they can be re-allocated very very quickly, because they are cheap to create and destroy
	*/

	go sayHello()                      // spins and runs sayHello in a "green Thread"
	time.Sleep(100 * time.Millisecond) // sleeping the main thread, so that we see the output from the go routine without using waitgroup

	var msg = "Hello from msg"

	wg.Add(1)
	go func() {
		// this function has access to msg, this is due to the concept of closures, but we created a dependency between the main and the anonymous function, so generally pass the dependency as value
		fmt.Println(msg)
		wg.Done()
	}()
	msg = "Goodbye" // race condition
	wg.Wait()

}

func sayHello() {
	fmt.Println("Hello")
}
