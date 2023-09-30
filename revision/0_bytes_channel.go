package main

import (
	"fmt"
	"sync"
	"time"
)

func isPrime(x *uint64) bool {
	for i := uint64(2); (i * i) <= *x; i++ {
		if *x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var x uint64 = 97
	p := &x

	s := fmt.Sprintf("%v isPrime?: %v", x, isPrime(p))
	fmt.Println(s)

	const N = 10
	var primes []uint64

	var wg sync.WaitGroup

	// is this compile time? no, golang doesn't support compile time
	go func() {
		// var primesArr = make([]uint64, 0, N)
		var primesArr []uint64

		wg.Add(1)
		for i := uint64(2); len(primesArr) < N; i++ {
			if isPrime(&i) {
				primesArr = append(primesArr, i)
			}
		}
		wg.Done()

		primes = primesArr
	}()

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan struct{}) // 🤯 0 bytes

	go func(done chan struct{}) {
		defer close(done)
		for {
			select {
			case <-done:
				{
					fmt.Println("Done")
					return
				}
			case tick := <-ticker.C:
				fmt.Println("Tick at", tick)
			}
		}
	}(done)

	time.Sleep(10 * time.Second)
	done <- struct{}{}

	ticker.Stop()

	wg.Wait()
	fmt.Println(primes)
}
