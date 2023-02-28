package main

import "fmt"

// go doesn't support pointer arithmetic, but we can use unsafe package for it
func main() {
	var a int = 42
	var b *int = &a

	fmt.Println(a, b)
	a = 27
	fmt.Println(a, *b)
	fmt.Printf("variable: %v pointer: %p\n", a, b)

	var ms *myStruct
	// fmt.Println(ms) // nil type
	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Println((*ms).foo)

	// arrow operator is used, just syntax sugar
	ms.foo = 42
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
