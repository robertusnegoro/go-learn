package main

import "testing"

func TestNewDec(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected dect length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card of Five of Clubs, but got %s", d[len(d)-1])
	}
}
