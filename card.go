package carddeck

import "fmt"

// Suit represents the suit of the card
type Suit uint8

const (
	// Spade suit
	Spade Suit = iota
	// Diamond suit
	Diamond
	// Club suit
	Club
	// Heart suit
	Heart
)

// Rank represents the value of card
type Rank uint8

const (
	// Ace value
	Ace Rank = iota + 1
	// Two value
	Two
	// Three value
	Three
	// Four value
	Four
	// Five value
	Five
	// Six value
	Six
	// Seven value
	Seven
	// Eight value
	Eight
	// Nine value
	Nine
	// Ten value
	Ten
	// Jack value
	Jack
	// Queen value
	Queen
	// King value
	King
)

// Card represents a single card of the deck with a rand and a suit
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %vs", c.Rank.String(), c.Suit.String())
}

// New creates a deck of card in new deck order and returns it
func New() []Card {
	deck := make([]Card, 52)
	var index uint8 = 0
	for i := 0; i < 4; i++ {
		for j := 1; j < 14; j++ {
			deck[index] = Card{
				Suit: Suit(i),
				Rank: Rank(j),
			}
			index++
		}
	}
	return deck
}
