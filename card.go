package deck

import (
	"fmt"
)

type Card struct {
	rank Rank
	suit Suit
}

func (c Card) String() string {
	return fmt.Sprintf("<Card: %s of %ss>", c.rank, c.suit)
}

func (c Card) Value() uint8 {
	return c.suit.index*13 + c.rank.index
}

func (c Card) Rank() Rank{
	return c.rank
}

func (c Card) Suit() Suit{
	return c.suit
}

