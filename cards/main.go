package main

// import "fmt"

func main() {
	cardsSlice := newDeck()
	// cardsSlice.print()

	// first, rest := deal(cardsSlice, 10)
	//
	// fmt.Println("first :")
	// first.print()
	//
	// fmt.Println("rest :")
	// rest.print()

	cardsSlice.createDeckFile()
}

func getCard() string {
	return "random card"
}
