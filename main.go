package main

func main() {
	// Declare a slice of strings
	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
