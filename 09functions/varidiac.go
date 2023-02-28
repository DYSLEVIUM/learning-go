package main

import "fmt"

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *s)
}

// varidac parameter has to be the last one
func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	// this isn't safe in other languages, as after this function finishes, this execution stack is destroyed, so we are returning a pointer to a location that just got freed
	return &result // but in go, it is automatically promoted to the shared memory(heap)
}
