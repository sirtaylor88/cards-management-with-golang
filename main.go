package main

import "fmt"

func main() {
	// Declare a slice of strings
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
