package main

func main() {

	cards := newDeck()

	cards = cards.shuffle()

	newDeck := cards.deal(4)

	newDeck.print()

}
