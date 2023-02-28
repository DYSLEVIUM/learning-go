package main

import "fmt"

func main() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	var ad func(int, int) int = func(a, b int) int {
		return a + b
	}

	s := ad(1, 2)
	fmt.Println("The sum is ", s)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}
