package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main(){
	// alex := person{"Alex","Anderson"}
	alex := person{ 
		firstName: "Alex", 
		lastName: "Anderson", 
		contactInfo: contactInfo{
			email: "Alex.Anderson@gmail.com",
			zipCode: 98027,
		},
	}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "Alex.Anderson@gmail.com"
	alex.print()
	alex.updateFirstName("jimmy")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerTpPerson *person) updateFirstName(newFirstName string) {
	(*pointerTpPerson).firstName = newFirstName
}