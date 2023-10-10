package main

var test int

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
	//fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
