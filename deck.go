package main

import "fmt"

// create new type of 'deck'

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	// for i, suit := range cardSuits { // i will not be used, so replace it with _ for every unused var
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// this is a receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this function returning multiple value
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
