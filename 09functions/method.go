package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()
	fmt.Println("The name is ", g.name)
}

type greeter struct {
	greeting string
	name     string
}

// we get copy of the object, change to pointer for the this object
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// func (g *greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// 	g.name = ""
// }
