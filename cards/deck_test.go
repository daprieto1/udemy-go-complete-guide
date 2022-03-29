package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(deck))
	}

	if deck[0] != "Hearts of Ace" {
		t.Errorf("Expected first card to be Hearts of Ace but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Clubes of Four" {
		t.Errorf("Expected last card to be Clubes of Four but got %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(deck))
	}

	os.Remove(filename)
}
