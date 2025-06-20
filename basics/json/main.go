package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{
		Name:    "Alice",
		Age:     30,
		IsAdult: true,
	}

	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("JSON output:", string(jsonData))

	var decodedPerson Person

	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}
	
	fmt.Println("Decoded Person:", decodedPerson)
}
