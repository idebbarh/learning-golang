package main

import (
	"fmt"
	"os"
)

type deck []string

func newDeck() deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	deckNames := []string{}
	for _, s := range suits {
		for _, r := range ranks {
			deckNames = append(deckNames, r+" of "+s)
		}
	}
	return deckNames
}

func (d deck) print() {
	for _, deck := range d {
		fmt.Println(deck)
	}
}

func deal(d deck, hand int) (deck, deck) {
	first := d[:hand]
	rest := d[hand:]

	return first, rest
}

func (d deck) createDeckFile() {
	filePath := "./deck.txt"

	deckFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	defer deckFile.Close()

	for _, l := range d {
		_, err = fmt.Fprintf(deckFile, l)

		if err != nil {
			return
		}

		_, err = fmt.Fprintf(deckFile, "\n")

		if err != nil {
			return
		}
	}

	fmt.Println("file create with data")
}
