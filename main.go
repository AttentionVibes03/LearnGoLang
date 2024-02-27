package main

var test int

func main() {
	cards := newDeck()


	hand, remainingCards := deal(cards, 5)
	//cards.print()
	//fmt.Println(cards)

	hand.print()
	remainingCards.print()
}

