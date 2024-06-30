package main

import "fmt"

// Create new type of 'deck'
// which is a slice of strings
type deck []string


func newDeck() deck{
	
	cards := deck{}
	
	cardSuits:= []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, suit + " of " + value)
		}
	}
	return cards
}

//"d" is similar to the OO "this"
func (d deck) print(){
	for i, card := range d{
		fmt.Println(i, card)
	}
}




