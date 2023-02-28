package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	// statePopulations := map[string]int{
	// 	"California": 314124,
	// 	"Texas":      241421,
	// 	"Florida":    21416161,
	// 	"Ohio":       2414,
	// }

	// fmt.Println(statePopulations)

	statePopulations := make(map[string]int, 10)
	statePopulations = map[string]int{
		"California": 314124,
		"Texas":      241421,
		"Florida":    21416161,
		"Ohio":       2414,
	}

	statePopulations["Georgia"] = 10125125
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	pop, ok := statePopulations["Georgia"]
	fmt.Println(pop, ok)

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Snow",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)

	bDoctor := struct{ name string }{name: "Cercy"} // anonymous struct
	fmt.Println(bDoctor)

	// composition (is-a)
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.Canfly = false

	fmt.Println(b)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

// composition
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float64
	Canfly   bool
}
