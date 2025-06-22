package main

import "fmt"

// Method - 1
// func displaySlice[T any](x []T) {
// 	for _, v := range x {
// 		fmt.Println(v)
// 	}
// }

// Method - 2
// func displaySlice[T comparable](x []T) {
// 	for _, v := range x {
// 		fmt.Println(v)
// 	}
// }

// Method - 3
// func displaySlice[T int | string](x []T) {
// 	for _, v := range x {
// 		fmt.Println(v)
// 	}
// }

type Stack[T any] struct {
	elements []T
}

func (s Stack[T]) displaySlice() {
	for _, v := range s.elements {
		fmt.Println(v)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"a", "b", "c", "d"}

	// displaySlice(nums)
	// displaySlice(strs)

	stack_num := Stack[int]{elements: nums}
	stack_num.displaySlice()

	stack_str := Stack[string]{elements: strs}
	stack_str.displaySlice()
}
