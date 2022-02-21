package main

func main() {
	// Declare a slice of strings
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
