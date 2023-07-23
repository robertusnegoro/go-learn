package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// tempVar := []string(d)
	// resultVar := strings.Join(tempVar, ",")
	// return resultVar
}

// set receiver function type of error because
// we don't need to think what is the writefile output
// instead, we output the error if it have problem
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1 : log the error and return a call to newDeck() X
		// option 2 : log the error and entirely quit the program :checked:
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
