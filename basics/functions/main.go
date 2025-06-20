package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function.")
}

func add(a, b int) int {
	return a + b
}

func sub(a int, b int) (result int) {
	result = a - b
	return
}

func div(a, b int) float64 {
	return float64(a) / float64(b)
}

func multiplyByTwo() func(a int) int {
	return func(a int) int {
		return a * 2
	}
}

func processFunction(fn func(a int) int) int {
	return fn(1)
}

func main() {
	simpleFunction()
	sum := add(5, 3)
	sub := sub(5, 3)
	div := div(10, 2)
	fmt.Println("The sum of 5 and 3 is:", sum)
	fmt.Println("The subtraction of 3 from 5 is:", sub)
	fmt.Println("The division of 5 by 2 is:", div)

	fn := multiplyByTwo()
	result := processFunction(fn)
	fmt.Println("The result of the function is:", result)
}
