package main

import "fmt"

//Create New Type Of Deck
// which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
