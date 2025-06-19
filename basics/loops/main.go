package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Number is:", i)
	}

	counter := 0
	for {
		fmt.Println("Counter is:", counter)
		counter++
		if counter >= 3 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, number := range numbers {
		fmt.Printf("Index: %d, Number: %d\n", index, number)
	}

	word := "Hello World"
	for index, char := range word {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}
