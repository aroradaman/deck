package deck

import (
	"fmt"
)

type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) String() string {
	return fmt.Sprintf("<Card: %s of %ss>", card.Rank, card.Suit)
}

func (card Card) Value() uint8 {
	return card.Suit.index*13 + card.Rank.index
}
