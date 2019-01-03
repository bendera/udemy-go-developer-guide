package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{"Ace of Spades", "Two of Spades"}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
