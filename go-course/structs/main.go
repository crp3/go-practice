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
	name := "Bill"
	namePointer := &name
	fmt.Println(namePointer)
	print(namePointer)

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func print(namePointer *string) {
	fmt.Println(namePointer)
}
