package main

// import "fmt"

func main() {
	cards := newDeck()
	// range is a default keyword
	// for i, card := range cards{
	// 	fmt.Println(i, card)
	// }
	// fmt.Println(cards)
	// fmt.Println(cards.toByteSlice())
	// cards := newDeckFromFile("HelloWorld")
	cards.shuffel()
	cards.print()
}
