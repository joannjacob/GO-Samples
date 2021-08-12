package main

import "fmt"

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
	// joann := person{firstName: "Joann", lastName: "Jacob"}
	// fmt.Println(joann)

	// var joann person

	// joann.firstName = "Joann"
	// joann.lastName = "Jacob"
	// fmt.Println(joann)
	// fmt.Printf("%+v", joann)

	joann := person{
		firstName: "Joann",
		lastName:  "Jacob",
		contactInfo: contactInfo{
			email:   "joann@gmail.com",
			zipCode: 1022,
		},
	}

	// joannptr := &joann
	joann.updateName("Josh")
	joann.print()
	// fmt.Printf("%+v", joann)

}

func (ptrToPerson *person) updateName(newFirstName string) {
	(*ptrToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
