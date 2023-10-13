package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type card struct {
	suit string
	rank string
}

type deck []card

func deckEqual(a, b deck) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (d deck) toString() string {
	joindedCards := ""
	for i, card := range d {
		fullCardName := card.rank + " of " + card.suit
		if i == 0 {
			joindedCards += fullCardName
		} else {
			joindedCards += "," + fullCardName
		}
	}
	return joindedCards
}

func toDeck(s string) deck {
	stringCards := strings.Split(s, ",")
	cardsDeck := deck{}

	for _, sc := range stringCards {
		sliceCard := strings.Split(sc, " of ")
		card := card{sliceCard[0], sliceCard[1]}
		cardsDeck = append(cardsDeck, card)
	}
	return cardsDeck
}

func newDeck() deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	deckNames := deck{}
	for _, s := range suits {
		for _, r := range ranks {
			card := card{s, r}
			deckNames = append(deckNames, card)
		}
	}
	return deckNames
}

func (d deck) print() {
	for _, deck := range d {
		fmt.Println(deck.suit + " of " + deck.rank)
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

	return toDeck(string(data)), err
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i int, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
