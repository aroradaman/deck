package deck

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck("standard")
	// 13 Ranks & 4 Suits
	if deck.Size() != 13*4 {
		t.Error(fmt.Sprintf("Expected 52 cards, got %d", deck.Size()))
	}
}

func TestDeck_Shuffle(t *testing.T) {
	deck := NewDeck("standard")
	oldDeck := NewDeck("standard")
	copy(deck.Cards, oldDeck.Cards)

	deck.Shuffle()

	if reflect.DeepEqual(deck.Cards, oldDeck.Cards) {
		t.Error("Expected different card ordering after shuffle, got same.")
	}

}

func TestDeck_DefaultSort(t *testing.T) {
	deck := NewDeck("standard")
	firstCard, lastCard := Card{rank: Two, suit: Heart}, Card{rank: Ace, suit: Spade}
	deck.Shuffle()

	// Swap first card with random card
	deck.Swap(0, 1+rand.Intn(51))

	// Swap last card with some random card
	deck.Swap(deck.Size()-1, 1+rand.Intn(51))

	// Sort the deck
	deck.DefaultSort()
	if !(deck.Cards[0] == firstCard && deck.Cards[deck.Size()-1] == lastCard) {
		t.Error(fmt.Sprintf("Expected first card after sorting %v, got %v", firstCard, deck.Cards[0]))
	}
}

func TestDeck_Sort(t *testing.T) {
	deck := NewDeck("standard")
	firstCard, lastCard := Card{rank: Two, suit: Heart}, Card{rank: Ace, suit: Spade}
	deck.Shuffle()

	// Swap first card with random card
	deck.Swap(0, 1+rand.Intn(51))

	// Swap last card with some random card
	deck.Swap(deck.Size()-1, 1+rand.Intn(51))

	less := func(i, j int) bool {
		return deck.Cards[i].Value() < deck.Cards[j].Value()
	}
	// Sort the deck using custom function
	deck.Sort(less)

	if !(deck.Cards[0] == firstCard && deck.Cards[deck.Size()-1] == lastCard) {
		t.Error(fmt.Sprintf("Expected first card after sorting %v, got %v", firstCard, deck.Cards[0]))
	}

}

func TestDeck_Filter(t *testing.T) {
	deck := NewDeck("standard")
	filter := func(card Card) bool {
		if card.suit == Heart || card.rank == Queen {
			return true
		}
		return false
	}

	deck.Filter(filter)

	for _, card := range deck.Cards {
		if card.suit == Heart || card.rank == Queen {
			continue
		} else {
			t.Error(fmt.Sprintf("Expected either Suit 'Heart' or Rank 'Queen', got %v", card))
		}
	}
}

func TestDeck_Add(t *testing.T) {
	deckA := NewDeck("standard")
	deckB := NewDeck("standard")

	deckA.Add(deckB)
	if deckA.Size() != 2*deckB.Size() {
		t.Error(fmt.Sprintf("Expected %d cards, got %d.", 2*deckB.Size(), deckA.Size()))
	}
}
