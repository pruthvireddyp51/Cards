package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

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

// Reciver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}

func (d deck) toByteSlice() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toByteSlice()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err!= nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffel() {
	for index := range d {
			newPosition := rand.Intn(len(d) - 1)
			d[index], d[newPosition] = d[newPosition], d[index]
	}
}