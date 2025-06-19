package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)

	var age int
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)
	fmt.Printf("You are %d years old.\n", age)
}
