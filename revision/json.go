package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// age  int    `json:"age"` // this will give error, as it is a json tag and has to be exported
}

func main() {
	// encoding
	encodedPerson := Person{Name: "Pushpa", Age: 23}
	data, err := json.Marshal(encodedPerson)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(data))

	// decoding
	var decodedPerson Person
	jsonStr := `{"name":"Pushpakant", "age":23}`
	err = json.Unmarshal([]byte(jsonStr), &decodedPerson)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(decodedPerson.Name, decodedPerson.Age)
}
