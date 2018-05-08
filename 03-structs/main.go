package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Setting up struct way 1
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }
	// fmt.Println(alex)

	// Setting up struct way 2
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmailz.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()

	// fmt.Printf("%+v", jim)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
