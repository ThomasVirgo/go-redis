package game

import (
	"fmt"
	"math/rand"
)

type Suit string
type Value string

const (
	DIAMONDS   Suit = "♦"
	HEARTS     Suit = "♥"
	SPADES     Suit = "♠"
	CLUBS      Suit = "♣"
	EMPTY_SUIT Suit = "empty"
)

const (
	ACE         Value = "A"
	TWO         Value = "2"
	THREE       Value = "3"
	FOUR        Value = "4"
	FIVE        Value = "5"
	SIX         Value = "6"
	SEVEN       Value = "7"
	EIGHT       Value = "8"
	NINE        Value = "9"
	TEN         Value = "10"
	JACK        Value = "11"
	QUEEN       Value = "12"
	KING        Value = "13"
	EMPTY_VALUE Value = "empty"
)

type Card struct {
	Suit   Suit
	Value  Value
	FaceUp bool
}

func (c *Card) ToString() string {
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}

var suits = []Suit{HEARTS, DIAMONDS, CLUBS, SPADES}
var values = []Value{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}

func NewDeck() []Card {
	var cards []Card
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, Card{Suit: suit, Value: value, FaceUp: false})
		}
	}

	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}
