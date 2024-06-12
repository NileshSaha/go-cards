package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	fmt.Println(cards)
	hand, remainingCards := deal(cards, 3)
	hand.print()
	remainingCards.print()

	fmt.Println(cards.toString())

	// greeting := "Hi There!"
	// fmt.Println([]byte(greeting))
}
