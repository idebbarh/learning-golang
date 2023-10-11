package main

import (
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := newDeck()
	newDeckL := len(newDeck)
	if newDeckL != 52 {
		t.Errorf("expected deck of length %d but got %d", 52, newDeckL)
	}
}

func TestDeal(t *testing.T) {
	newDeck := newDeck()
	for i := 0; i < 10; i++ {
		r := rand.Intn(len(newDeck))
		left, right := deal(newDeck, r)
		if len(left) != r {
			t.Errorf("expected left part of length %d but got %d", r, len(left))
		}
		if len(right) != len(newDeck)-r {
			t.Errorf("expected right part of length %d but got %d", len(newDeck)-r, len(right))
		}
	}
}

func TestSaveToFileAndGetFromFile(t *testing.T) {
	newDeck := newDeck()
	fileName := "test.txt"
	err := newDeck.saveToFile(fileName)
	if err != nil {
		t.Errorf("error on save to file function")
	}

	fileDeck, err := getFromFile(fileName)
	if err != nil {
		t.Errorf("error on get from file function")
	}

	if reflect.TypeOf(fileDeck) != reflect.TypeOf(newDeck) {
		t.Errorf("expected getFromFile function to return type of %v but got type of %v", reflect.TypeOf(newDeck), reflect.TypeOf(fileDeck))
	} else if !deckEqual(newDeck, fileDeck) {
		t.Errorf("create deck and the deck from the file are not equal")
	}

	err = os.Remove(fileName)

	if err != nil {
		t.Errorf("Error removing temporary file: %v", err)
	}
}

func TestShuffle(t *testing.T) {
	deckToShuffle := newDeck()
	deckToCompare := newDeck()

	deckToShuffle.shuffle()

	for i := range deckToShuffle {
		if deckToShuffle[i] != deckToCompare[i] {
			break
		}

		if i == len(deckToShuffle)-1 {
			t.Errorf("expected different deck after shuffle but got the same")
		}
	}
}
