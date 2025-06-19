package main

import (
	"fmt"
)

func main() {
	var message string = "Hello, World!"
	var version = 1.0
	fmt.Println(message)
	fmt.Println("Version:", version)

	var money int = -100
	fmt.Println("Money:", money)

	var decided = true
	fmt.Println("Decided:", decided)

	const pi = 3.14
	fmt.Println("Pi:", pi)

	name := "John Doe"
	fmt.Println("Name:", name)

	// Variables can be exported if they start with a capital letter
	Name := "Jane Doe"
	fmt.Println("Name with capital N:", Name)
}
