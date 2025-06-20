package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0

	for i, num := range nums {
		sum += num
		fmt.Printf("Index: %d, Value: %d, Running Sum: %d\n", i, num, sum)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	for i, c := range "hello" {
		fmt.Println("Index:", i, "Character:", string(c))
	}
}
