package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	fmt.Println(cards)
}

func newCard() string {
	return "Five of diamonds"
}
