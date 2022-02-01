package main

func main() {
	cards := newDeckFromFile("myassisonfire.txt")
	cards.shuffle()

	hand, remainingDeck := deal(cards, 15)
	hand.print()
	remainingDeck.toString()
	hand.saveToFile("handSave.txt")
	cards.print()
}
