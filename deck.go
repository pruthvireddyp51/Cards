package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"h1", "h2", "h3"}
	cardValues := []string{"v1", "v2", "v3"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Reciver on a function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
