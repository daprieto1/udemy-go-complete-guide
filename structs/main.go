package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Paredes",
		contact: ContactInfo{
			email:   "jim@mail.com",
			zipCode: 21345,
		},
	}

	jim.update("Jimmy")
	jim.print()
}

func (p *Person) update(firstName string) {
	(*p).firstName = firstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
