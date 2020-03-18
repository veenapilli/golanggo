package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected: 16. Got: %v", len(d))
	}
}

func TestAssertDeck(t *testing.T) {
	d := newDeck()
	if d[0] != "a of spades" {
		t.Errorf("Expected: a of spades. Got: %v", d[0])
	}
	if d[15] != "4 of clover" {
		t.Errorf("Expected: 4 of clover. Got: %v", d[15])
	}
}

func TestSavingAndReadingDeck(t *testing.T) {
	fileName := "_testDeck.txt"

	os.RemoveAll(fileName)

	d := newDeck()

	writeDeck(d, fileName)

	r := readDeck(fileName)
	if len(r) != 16 {
		t.Errorf("Expected: 16 cards. Got: %v", len(r))
	}

}
