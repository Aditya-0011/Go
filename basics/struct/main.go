package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	Street string
	City   string
	State  string
}

type Employee struct {
	Person  Person
	Contact Contact
	Address Address
}

func main() {
	var man Person
	man.FirstName = "John"
	man.LastName = "Doe"
	man.Age = 30
	fmt.Println("Man:", man)

	woman := Person{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       28,
	}
	fmt.Println("Woman:", woman)

	var child = new(Person)
	child.FirstName = "Alice"
	child.LastName = "Doe"
	child.Age = 5
	fmt.Println("Child:", *child)

	employee := Employee{
		Person: Person{
			FirstName: "Bob",
			LastName:  "Smith",
			Age:       40,
		},
		Contact: Contact{
			Email: "a@a.a",
			Phone: "123-456-7890",
		},
		Address: Address{
			Street: "123 Main Street",
			City:   "Anytown",
			State:  "CA",
		},
	}
	fmt.Println("Employee details:", employee)
	fmt.Println("Person:", employee.Person)
	fmt.Println("Contact:", employee.Contact)
	fmt.Println("Address:", employee.Address)
}
