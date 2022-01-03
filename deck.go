package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Deck struct {
	name  string
	cards []Card
}

func NewDeck(name string) Deck {
	var cards []Card
	for _, suit := range SUITS {
		for _, rank := range RANKS {
			cards = append(cards, Card{suit: suit, rank: rank})
		}
	}
	return Deck{name: name, cards: cards}
}

func (deck Deck) Swap(i, j int) {
	deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
}

func (deck Deck) Size() int {
	return len(deck.cards)
}

func (deck Deck) Less(i, j int) bool {
	return deck.cards[i].Value() < deck.cards[j].Value()
}

func (deck Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(deck.Size(), deck.Swap)
}

func (deck Deck) DefaultSort() {
	sort.Slice(deck.cards, deck.Less)
}

func (deck Deck) Sort(less func (i, j int) bool){
	sort.Slice(deck.cards, less)
}

func (deck Deck) String() string {
	return fmt.Sprintf("Deck %s; Fir" +
		"st Card: %v; Last Card: %v", deck.name, deck.cards[0], deck.cards[deck.Size()-1])
}

func (deck *Deck) Add(newDeck Deck) {
	deck.cards = append(deck.cards, newDeck.cards...)
}

func (deck *Deck) Filter(filter func (card Card) bool) {
	cards := make([]Card, 0)
	for _, c := range deck.cards{
		if filter(c){
			cards = append(cards, c)
		}
	}
	deck.cards = cards
}