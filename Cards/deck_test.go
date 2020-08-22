package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Error Length %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Error Card %v", d[0])
	}
	if d[2] != "Three of Spades" {
		t.Errorf("Error Card %v", d[3])
	}
}
