package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type germanBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	gb := germanBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(gb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting1(sb spanishBot) {
// 	fmt.Println(sb.getGreeting1())
// }

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (germanBot) getGreeting() string {
	return "Guten tag"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
