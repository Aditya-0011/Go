package main

import "fmt"

func main() {

	var name [5]string
	name[0] = "Alice"
	name[1] = "Bob"
	name[2] = "Charlie"
	name[3] = "Diana"
	name[4] = "Eve"
	fmt.Println("Names in the array:", name)
	fmt.Println("Length of the names array:", len(name))

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers in the array:", numbers)

	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("2D Array:", nums)
}
