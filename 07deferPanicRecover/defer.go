package main

import "fmt"

// defer is commonly used to open a resource and immediately close a resource, but if there are many resources then we may want to manually close a resource
func main() {
	fmt.Println("start")
	defer fmt.Println("middle") // defered functions are called in lifo order after the function exits
	fmt.Println("end")
}
