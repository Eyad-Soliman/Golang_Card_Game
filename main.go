package main

import "fmt"

func main() {

	//creates a deck
	deck := newDeck()

	//shuffles the deck
	deck.shuffle()

	//Creates 4 hands of 4 cards each
	hand1, remainingDeck := deal(deck, 4)
	hand2, remainingDeck := deal(remainingDeck, 4)
	hand3, remainingDeck := deal(remainingDeck, 4)
	hand4, remainingDeck := deal(remainingDeck, 4)

	//Displays hands and remaining deck
	fmt.Println("Hand 1 contains: ")
	hand1.print()
	fmt.Println("Hand 2 contains: ")
	hand2.print()
	fmt.Println("Hand 3 contains: ")
	hand3.print()
	fmt.Println("Hand 4 contains: ")
	hand4.print()
	fmt.Println("Cards remaining in deck: ")
	remainingDeck.print()
}
