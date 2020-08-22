package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	house := deck{"Spades", "Hearts", "Clubs", "Diamonds"}
	rank := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}
	for i := range house {
		for j := range rank {
			cards = append(cards, rank[j]+" of "+house[i])
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}

}

func (d deck) deal(n int) deck {
	//fmt.Println(n)
	newDeck := deck{}
	r := rand.Intn(len(d) - n)
	newDeck = append(newDeck, d[r:r+n]...)
	d = append(d[:r], d[r+n:]...)
	return newDeck
}

func (d deck) shuffle() deck {
	s := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(s)
	n := rn.Intn(len(d) - 1)
	for i := 0; i < n; i++ {
		r := rn.Intn(len(d) - 1)
		d[i], d[r] = d[r], d[i]
	}
	return d
}
