package main

import "fmt"

func main() {
	studentMarks := make(map[string]int)

	studentMarks["Alice"] = 90
	studentMarks["Bob"] = 85
	studentMarks["Charlie"] = 92
	studentMarks["David"] = 88

	for student, marks := range studentMarks {
		fmt.Printf("Student: %s, Marks: %d\n", student, marks)
	}

	delete(studentMarks, "Bob")
	for student, marks := range studentMarks {
		fmt.Printf("Student: %s, Marks: %d\n", student, marks)
	}

	marks, exists := studentMarks["Bob"]
	fmt.Printf("Bob's marks: %d, Bob Exists?: %t\n", marks, exists)

	person := map[string]string{
		"name":    "Alice",
		"age":     "30",
		"country": "USA",
	}

	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

}
