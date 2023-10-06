package main

import "fmt"

func main() {
	cardsSlice := newDeck()

	cardsSlice = cardsSlice.shuffle()
	fmt.Println(len(cardsSlice))
}
