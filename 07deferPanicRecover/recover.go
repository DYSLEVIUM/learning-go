package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End") // we recovered from the panic and exectued the functions higher-up the call stack
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// panic(err) // we can panic again, if we can't handle the panic
		}
	}()
	panic("Something Bad Happened")
	fmt.Println("Done Panicking")
}
