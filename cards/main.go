package main

func main() {
	cardsSlice := newDeck()
	cardsSlice.saveToFile("cards.txt")
	loadedCards, _ := getFromFile("cards.txt")
	loadedCards.shuffle()
	loadedCards.print()
}
