package main

func main() {
	cardsSlice := newDeck()
	cardsSlice.shuffle()
	cardsSlice.print()
}
