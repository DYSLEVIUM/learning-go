package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("==")

	for i := 0; i < 5; i = i + 2 {
		fmt.Println(i)
	}
	fmt.Println("==")

	for i, j := 0, 0; i < 5; i, j = i+2, j+1 {
		fmt.Println(i)
		fmt.Println(j)
	}
	fmt.Println("==")

	i := 0
	for {
		fmt.Println(i)
		i++
		if i < 10 {
			continue
		} else {
			break
		}
	}
	fmt.Println("==")

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
	fmt.Println("==")
}
