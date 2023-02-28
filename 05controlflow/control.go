package main

import "fmt"

func main() {
	if true {
		fmt.Println("The Test is true")
	}

	statePopulations := map[string]int{
		"California": 314124,
		"Texas":      241421,
		"Florida":    21416161,
		"Ohio":       2414,
	}

	// initializer syntax
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	fmt.Println(returnTrue())

	switch 2 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	i := 10
	switch {
	case i < 10:
		fmt.Println("less than 10")
		fallthrough // default is implied, so we use fallthrough
	case i < 20:
		fmt.Println("less than 20")
	default:
		fmt.Println("def")
	}

	var t interface{} = [3]int{}
	switch t.(type) {
	case int:
		fmt.Println("t is int")
		break
		fmt.Println("This will not print")
	case string:
		fmt.Println("t is string")
	case float64:
		fmt.Println("t is float64")
	case [3]int:
		fmt.Println("t is [3]int")
	default:
		fmt.Println("t is another type")
	}
}

func returnTrue() bool {
	fmt.Println("returning True")
	return true
}
