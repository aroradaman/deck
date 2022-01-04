package deck

type Rank struct {
	index uint8
	name  string
}

var (
	Two   = Rank{index: 2, name: "Two"}
	Three = Rank{index: 3, name: "Three"}
	Four  = Rank{index: 4, name: "Four"}
	Five  = Rank{index: 5, name: "Five"}
	Six   = Rank{index: 6, name: "Six"}
	Seven = Rank{index: 7, name: "Seven"}
	Eight = Rank{index: 8, name: "Eight"}
	Nine  = Rank{index: 9, name: "Nine"}
	Ten   = Rank{index: 10, name: "Ten"}
	Jack  = Rank{index: 11, name: "Jack"}
	Queen = Rank{index: 12, name: "Queen"}
	King  = Rank{index: 13, name: "King"}
	Ace   = Rank{index: 14, name: "Ace"}
)

var RANKS = [...]Rank{Ace, King, Queen, Jack, Ten, Nine, Eight, Seven, Six, Five, Four, Three, Two}

func (r Rank) String() string {
	return r.name
}

func (r Rank) Value()int{
	return int(r.index)
}
