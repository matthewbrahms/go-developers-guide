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
	contact   contactInfo
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
		contact: contactInfo{
			email:   "jim@gmailz.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
