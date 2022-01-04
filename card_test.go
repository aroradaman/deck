package deck

import (
	"fmt"
)

func ExampleCard() {
	fmt.Println(Card{rank: Ace, suit: Spade})
	fmt.Println(Card{rank: Two, suit: Heart})
	fmt.Println(Card{rank: King, suit: Club})

	// Output:
	// <Card: Ace of Spades>
	// <Card: Two of Hearts>
	// <Card: King of Clubs>
}
