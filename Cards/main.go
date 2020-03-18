package main

import "fmt"

var card string
var fileName string = "deal.txt"

func main() {
	cards := newDeck()
	//cards.print()

	deal, _ := deal(cards, 3)
	fmt.Println("Deal")
	deal.print()
	// fmt.Println("Remaining")
	// remaining.print()

	writeDeck(deal, fileName)
	cards = readDeck(fileName)
	cards.print()

	fmt.Println("Shuffle")
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "123"
}
