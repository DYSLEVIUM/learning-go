package main

import (
	"fmt"
	"strconv"
)

var name string = "Pushpakant Behera"

var (
	fname string = "Pushpa"
	lname string = "Behera"
)

func main() {
	var i int
	i = 42
	i = 27

	var j int = 10

	k := 30

	var l float32 = 22

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(name)
	fmt.Println(fname)
	fmt.Println(lname)

	fmt.Printf("%v, %T\n", j, j)

	// strings are immutable
	var ss string
	ss = strconv.Itoa(i)

	fmt.Printf("%v, %T\n", ss, ss)

	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	var t uint = 10
	fmt.Println(t)

	// and-not -> a:= 10, b:= 3 => a&^b => not of or
	a := 10
	b := 3
	fmt.Println(a &^ b)

	var x float64 = 214.234e21
	fmt.Println(x)

	var z complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", z, z)

	fmt.Println(real(z))
	fmt.Println(imag(z))

	var z2 complex128 = complex(10, 14)
	fmt.Printf("%v, %T\n", z2, z2)
}
