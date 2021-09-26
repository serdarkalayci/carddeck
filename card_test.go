package carddeck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardStringFromInt(t *testing.T) {
	card := Card{
		Suit: 0,
		Rank: 11,
	}
	assert.Equal(t, "Jack of Spades", card.String())
}

func TestCardStringFromString(t *testing.T) {
	card := Card{
		Suit: Heart,
		Rank: Queen,
	}
	assert.Equal(t, "Queen of Hearts", card.String())
}

func TestCardStringFromIntErrSuit(t *testing.T) {
	card := Card{
		Suit: 4,
		Rank: 11,
	}
	assert.Equal(t, "Jack of Suit(4)s", card.String())
}

func TestCardStringFromIntErrRank(t *testing.T) {
	card := Card{
		Suit: 2,
		Rank: 16,
	}
	assert.Equal(t, "Rank(16) of Clubs", card.String())
}

func TestNewDeck(t *testing.T) {
	deck := New()
	assert.Equal(t, "Three of Spades", deck[2].String())
	assert.Equal(t, "Queen of Diamonds", deck[24].String())
	assert.Equal(t, "King of Hearts", deck[51].String())
}
