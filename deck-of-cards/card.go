//go:generate stringer -type=Suite,Face
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suite int
type Face int

const (
	Spades Suite = iota + 1
	Diamonds
	Clubs
	Hearts
	Joker
)

var suits = [...]Suite{Spades, Diamonds, Clubs, Hearts}

const (
	Ace Face = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	firstCard = Ace
	lastCard  = King
)

type Card struct {
	Suite
	Face
}

func (c Card) String() string {
	if c.Suite == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %s", c.Face.getFaceCard(), c.Suite.getSuiteName())
}

func (s Suite) getSuiteName() string {
	if s == 5 {
		return "Joker"
	} else {
		return [...]string{"Spades", "Diamonds", "Clubs", "Hearts"}[s-1]
	}

}

func (f Face) getFaceCard() string {
	return [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}[f-1]
}

func Shuffle(cards []Card) []Card {
	var shuffledCards []Card
	//This is added code, but if we want randomness everytime a new iteration is run, we can add some randomness to the source
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, value := range r.Perm(len(cards)) {
		shuffledCards = append(shuffledCards, cards[value])
	}

	return shuffledCards
}

// This opts ...func([]Card) []Card is a first order function, where functions are passed as values
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for i := firstCard; i <= lastCard; i++ {
			cards = append(cards, Card{Face: i, Suite: suit})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// This is the basic sorting function that have the less function which take two items from a slice and sort them
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Now if we want to have different type of sorting we needed to have a custom less function
// So this sort function takes the less as a parameter and returns the default sort as a function so that we can input in the new function
// This option function is like this New(Sort(Less(cards)))
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// In js, we use callback function right, this is the callback function that just takes out the absolute value and sort them
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// Now we can use this opts to find the absolute rank of the card
func absRank(c Card) int {
	return (int(c.Suite)-1)*int(lastCard) + int(c.Face)
}

// We will add one more option function Jokers, to add the number of jokers required in the function
// Something like New(Jokers(3)) to add 3 jokers
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suite: Joker, Face: Face(i)})
		}
		return cards
	}
}

// We will now add one more function to remove/filter the unwanted cards
// Another option function for New
func FilterCards(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var deckOfCards []Card
		for _, c := range cards {
			if !f(c) {
				deckOfCards = append(deckOfCards, c)
			}
		}
		return deckOfCards
	}
}

// Now for blackjack we should have an option so that we can enter multiple decks
// Another option function for New
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var newDeck []Card
		for i := 1; i <= n; i++ {
			newDeck = append(newDeck, cards...)
		}
		return newDeck
	}
}
