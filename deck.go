package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Deck struct {
	Name  string
	Cards []Card
}

func NewDeck(name string) Deck {
	var cards []Card
	for _, suit := range SUITS {
		for _, rank := range RANKS {
			cards = append(cards, Card{suit: suit, rank: rank})
		}
	}
	return Deck{Name: name, Cards: cards}
}

func (deck Deck) Swap(i, j int) {
	deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
}

func (deck Deck) Size() int {
	return len(deck.Cards)
}

func (deck Deck) Less(i, j int) bool {
	return deck.Cards[i].Value() < deck.Cards[j].Value()
}

func (deck Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(deck.Size(), deck.Swap)
}

func (deck Deck) DefaultSort() {
	sort.Slice(deck.Cards, deck.Less)
}

func (deck Deck) Sort(less func(i, j int) bool) {
	sort.Slice(deck.Cards, less)
}

func (deck Deck) String() string {
	return fmt.Sprintf("<Deck: %s; Fir"+
		"st Card: %v; Last Card: %v>", deck.Name, deck.Cards[0], deck.Cards[deck.Size()-1])
}

func (deck *Deck) Add(newDeck Deck) {
	deck.Cards = append(deck.Cards, newDeck.Cards...)
}

func (deck *Deck) Filter(filter func(card Card) bool) {
	cards := make([]Card, 0)
	for _, c := range deck.Cards {
		if filter(c) {
			cards = append(cards, c)
		}
	}
	deck.Cards = cards
}

func (deck *Deck) Draw() Card {
	card := deck.Cards[deck.Size()-1]
	deck.Cards = deck.Cards[:deck.Size()-1]
	return card
}

func (deck Deck) IsEmpty() bool {
	if len(deck.Cards) == 0 {
		return true
	}
	return false
}
