package main

import "fmt"

func main() {
	sayGreeting("Hello", "Pushpa")
}

// func sayGreeting(greeting string, name string) {
// 	fmt.Println(greeting, name)
// }

// same as above
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}
