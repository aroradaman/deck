package deck

import (
	"fmt"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Spade})
	fmt.Println(Card{Rank: Two, Suit: Heart})
	fmt.Println(Card{Rank: King, Suit: Club})

	// Output:
	// <Card: Ace of Spades>
	// <Card: Two of Hearts>
	// <Card: King of Clubs>
}
