package main

// import "fmt"

func main() {
	cards := deck{"Hello", cardName()}
	cards = newDeck()
	// range is a default keyword
	// for i, card := range cards{
	// 	fmt.Println(i, card)
	// }
	// fmt.Println(cards)
	cards.print()
}

func cardName() string {
	return "Name of the card"
}
