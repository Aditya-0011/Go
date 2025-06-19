package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,orange"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str2 := "   Hello,     World!     "
	trimmed := strings.TrimSpace(str2)
	fmt.Println(trimmed)

	firstname := "John"
	lastname := "Doe"
	fullName := strings.Join([]string{firstname, lastname}, " ")
	fmt.Println(fullName)
}
