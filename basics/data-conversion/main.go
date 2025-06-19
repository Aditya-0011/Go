package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Printf("num data-type: %T\n", num)

	var data float64 = float64(num)
	fmt.Printf("data data-type: %T\n", data)

	var str string = strconv.Itoa(num)
	fmt.Printf("str data-type: %T\n", str)

	var number string = "123"
	var convertedNum, _ = strconv.Atoi(number)
	fmt.Printf("convertedNum data-type: %T\n", convertedNum)

	var decimal string = "3.14"
	var floatNum, _ = strconv.ParseFloat(decimal, 64)
	fmt.Printf("floatNum data-type: %T\n", floatNum)
}
