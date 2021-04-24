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

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

func main() {

	jin := person{
		firstName: "Jin",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "example@gmail.com",
			zipCode: 94000,
		},
	}

	jin.updateName("jimmy")
	jin.print()

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
