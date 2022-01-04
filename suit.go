package deck

type Suit struct {
	index uint8
	name  string
}

var (
	Spade   = Suit{index: 3, name: "Spade"}
	Club    = Suit{index: 2, name: "Club"}
	Diamond = Suit{index: 1, name: "Diamond"}
	Heart   = Suit{index: 0, name: "Heart"}
)

var SUITS = [...]Suit{Spade, Club, Diamond, Heart}

func (s Suit) String() string {
	return s.name
}

func (s Suit) Value() int {
	return int(s.index)
}
