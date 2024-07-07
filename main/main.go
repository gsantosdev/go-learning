package main

import (
	"fmt"
)

func main() {
	//cards - part 1
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("cards")
	cardFromfile := newDeckFromFile("cards")

	cardFromfile.shuffle()
	cardFromfile.print()

	//person - part 2
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var alex2 person
	fmt.Println(alex2)

	alex2.firstName = "Alex"
	alex2.lastName = "Anderson"
	fmt.Println(alex2)

	//without pointer - update with pointer
	alex.print()
	alex.updateName("Alexy")
	alex.print()

	alexPointer := &alex
	alexPointer.updateName("Alexander")
	alex.print()

}
