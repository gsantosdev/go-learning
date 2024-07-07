package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
