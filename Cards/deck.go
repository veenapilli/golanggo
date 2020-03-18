package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// new type deck
// slice of strings

type deck []string

/**
Creates and returns a new deck of cards
*/
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spades", "diamonds", "hearts", "clover"}
	cardValues := []string{"a", "2", "3", "4"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

/**
Provides the size of the deck
*/
func (d deck) size() int {
	return len([]string(d))
}

/**
deal i cards form the deck
*/
func deal(d deck, i int) (deck, deck) {
	return d[:i], d[i:]
}

/**
@param d: deck to be written to the file
@param s: filename
*/
func writeDeck(d deck, s string) error {
	var deck string
	deck = d.toString()
	err := ioutil.WriteFile(s, []byte(deck), 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return err
}

/**
@param s: filename to read the deck from
*/
func readDeck(s string) deck {
	content, err := ioutil.ReadFile(s)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Printf("File contents: %s", content)
	fmt.Println()
	conString := string(content)
	conStrings := strings.Split(conString, ",")
	cards := deck(conStrings)
	return cards

}

/**
Shuffle cards
For each card, generate a random number 0 to n-1 and swap card
*/

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		swap := r.Intn(d.size())
		//d[i] = d[swap]
		//d[swap] = card
		d[i], d[swap] = d[swap], d[i]
	}
	return d
}

/**
receiver variable is 1 or 2 letters of the type
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/**
converts a deck to a string
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
