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

func (d Deck) Sort(less func(i, j int) bool) {
	sort.Slice(d.Cards, less)
}

func (d Deck) String() string {
	return fmt.Sprintf("<Deck: %s; Fir"+
		"st Card: %v; Last Card: %v>", d.Name, d.Cards[0], d.Cards[d.Size()-1])
}

func (d *Deck) Add(newDeck Deck) {
	d.Cards = append(d.Cards, newDeck.Cards...)
}

func (d *Deck) Filter(filter func(card Card) bool) {
	cards := make([]Card, 0)
	for _, c := range d.Cards {
		if filter(c) {
			cards = append(cards, c)
		}
	}
	d.Cards = cards
}

func (d *Deck) Draw() Card {
	card := d.Cards[d.Size()-1]
	d.Cards = d.Cards[:d.Size()-1]
	return card
}

func (d Deck) IsEmpty() bool {
	if len(d.Cards) == 0 {
		return true
	}
	return false
}
