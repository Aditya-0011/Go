package main

import "fmt"

func main() {
	x := 0

	if x < 0 {
		fmt.Println("x is less than 0")
	} else if (x >= 0) && (x < 5) {
		fmt.Println("x is less than 5 and greater than or equal to 0")
	} else {
		fmt.Println("x is greater than or equal to 5")
	}
}
