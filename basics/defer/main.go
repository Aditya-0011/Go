package main

import "fmt"

//defer creates a stack of function calls that are executed in reverse order after the surrounding function returns.

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting the program...")
	data := add(1, 2)
	defer fmt.Println("This is printed last.")
	defer fmt.Println("The result is:", data)
	fmt.Println("Doing some work...")
	fmt.Println("Work done!")
}
