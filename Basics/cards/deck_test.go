package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, instead got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, instead got %v", d[0])
	}

	if d[len(d)-1] != "King of Club" {
		t.Errorf("Expected last card King of Club, instead got %v", d[0])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove(".//test//*")

	deck := newDeck()
	deck.saveToFile(".//test//deckToFile.txt")

	savedDeck := newDeckFromFile(".//test//deckToFile.txt")

	if len(savedDeck) != len(deck) {
		t.Errorf("Expected %v cards in the deck found %v", len(deck), len(savedDeck))
	}
	os.Remove(".//test//*")
}
