package main

import (
	"fmt"
)

// race condition, either the go routine will finish first and hence the println, else if the main function finishes first, no message is printed
func main() {
	ch := make(chan string)

	go func() {
		msg := <-ch
		defer close(ch)

		fmt.Println(msg)
	}()

	ch <- "hello"
}
