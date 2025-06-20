package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	month := "March"
	switch month {
	case "January", "February", "March":
		fmt.Println("First quarter of the year")
	case "April", "May", "June":
		fmt.Println("Second quarter of the year")
	case "July", "August", "September":
		fmt.Println("Third quarter of the year")
	case "October", "November", "December":
		fmt.Println("Fourth quarter of the year")
	default:
		fmt.Println("Invalid month")
	}

	temperature := 30
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 20:
		fmt.Println("Cold")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	case temperature >= 30:
		fmt.Println("Hot")
	default:
		fmt.Println("Invalid temperature")
	}

	whoAmI := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a boolean")
		default:
			fmt.Println("other type")
		}
	}

	whoAmI(42)
	whoAmI("Hello")
	whoAmI(true)
	whoAmI(3.14)
}
