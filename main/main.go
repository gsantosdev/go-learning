package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("cards")
	cardFromfile := newDeckFromFile("cards")

	cardFromfile.shuffle()
	cardFromfile.print()
}
