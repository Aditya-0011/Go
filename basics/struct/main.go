package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) changeAge(age int) {
	p.Age = age
}

func (p Person) fullName() string {
	return p.FirstName + " " + p.LastName
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

func newAddress(street, city, state string) *Address {
	address := Address{
		Street: street,
		City:   city,
		State:  state,
	}

	return &address
}

type Employee struct {
	Person
	Contact
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

	employee.Person.changeAge(45)
	fmt.Println("Updated Employee Age:", employee.Person.Age)
	fmt.Println("Full Name:", employee.Person.fullName())

	address := newAddress("456 Elm Street", "Othertown", "NY")
	fmt.Println("New Address:", *address)

	country := struct {
		name    string
		capital string
	}{"Wonderland", "Wonder City"}
	fmt.Println("Country:", country)
}
