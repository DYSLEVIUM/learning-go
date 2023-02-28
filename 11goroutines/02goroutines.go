package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

// race conditions
// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	fmt.Printf("Hello #%v\n", counter)
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	wg.Done()
// }

// var m = sync.RWMutex{} // lock

// func main() {
// 	runtime.GOMAXPROCS(100)
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	m.RLock()
// 	fmt.Printf("Hello #%v\n", counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func increment() {
// 	m.Lock()
// 	counter++
// 	m.Unlock()
// 	wg.Done()
// }

var m = sync.RWMutex{} // lock

func main() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	runtime.GOMAXPROCS(100) // tuning parameter
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
