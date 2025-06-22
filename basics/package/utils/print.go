package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintMessage(message ...string) {
	for i, msg := range message {
		if i == 0 && len(message) > 1 {
			fmt.Println()
		}
		fmt.Printf("%s", msg)
	}

	//fmt.Println("\nAnother Random Number:", RandomInt(10))
	random := fmt.Sprintf("\nAnother Random number: %d", RandomInt(15))
	color.Red(random)
}
