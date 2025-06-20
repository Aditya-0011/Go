package main

import (
	"fmt"
	"slices"
)

func main() {
	number := []int{1, 2, 3, 4, 5}

	number = append(number, 6)
	fmt.Println("Numbers:", number)
	fmt.Println("Numbers Length:", len(number))
	fmt.Println("Numbers Capacity:", cap(number))

	name := []string{}

	name = append(name, "Alice", "Bob", "Charlie", "Diana", "Eve")
	fmt.Println("Names:", name)
	fmt.Println("Names Length:", len(name))
	fmt.Println("Names Capacity:", cap(name))

	animals := make([]string, 3, 5)
	animals = append(animals, "Dog", "Cat", "Elephant")
	fmt.Println("Animals:", animals)
	fmt.Println("Animals Length:", len(animals))
	fmt.Println("Animals Capacity:", cap(animals))

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println("Slice a:", a)
	fmt.Println("Slice b:", b)

	c := []int{1, 2, 3, 4, 5}
	fmt.Println("Slic c is same as a?", slices.Equal(c, a))
}
