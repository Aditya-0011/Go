package main

import "fmt"

func multiplyByTwo(num *int) {
	*num = *num * 2
}

func main() {
	// Method - 1
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	// Method - 2
	num := 2
	ptr := &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value at address of num:", *ptr)
	multiplyByTwo(&num)
	fmt.Println("Value after multiplying by two:", *ptr)
}
