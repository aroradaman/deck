package deck

import (
	"fmt"
)

func ExampleCard() {
	fmt.Println(Card{rank: Ace, suit: Spade})
	fmt.Println(Card{rank: Two, suit: Heart})
	fmt.Println(Card{rank: King, suit: Club})

	// Output:
	// Ace of Spades
	// Two of Hearts
	// King of Clubs
}
