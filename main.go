package main

func main() {
	cards := newDeckFromFile("my_cards.txt")
	cards.print()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// fmt.Println("----")
	// remainingDeck.print()
}
