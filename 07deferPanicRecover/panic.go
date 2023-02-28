package main

import "fmt"

// panic happen after the defered statements
func main() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

	// panic("Something Bad Happened")
}
