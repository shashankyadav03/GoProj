package main

import "fmt"

var car int

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Hearts"
	fmt.Println(card)
	card = newCard()
	fmt.Println(card)

	car = 9
	fmt.Println(car)

	cards := []string{newCard(), "Ace"}
	fmt.Println(cards)
	cards = append(cards, "Six")
	fmt.Println(cards)

	for i := range cards {
		fmt.Println(cards[i])
	}

}

func newCard() string {
	return "five"
}
