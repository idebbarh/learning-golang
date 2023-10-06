package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) toString() string {
	return strings.Join(d, ",")
}

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

func (d deck) saveToFile(filename string) error {
	deckString := d.toString()

	return os.WriteFile(filename, []byte(deckString), 6642)
}

func getFromFile(filename string) (deck, error) {
	data, err := os.ReadFile(filename)

	return strings.Split(string(data), ","), err
}

func (d deck) shuffle() {
	for i, v := range d {
		randomIndx := rand.Intn(len(d))
		tmp := d[randomIndx]
		d[randomIndx] = v
		d[i] = tmp
	}
}
