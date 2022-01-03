package deck

import (
	"fmt"
)

type Card struct {
	rank Rank
	suit Suit
}

func (card Card) String() string {
	return fmt.Sprintf("%s of %ss", card.rank, card.suit)
}

func (card Card) Value() uint8 {
	return card.suit.index*13 + card.rank.index
}
