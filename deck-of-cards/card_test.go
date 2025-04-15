package deck

import (
	"fmt"
	"testing"
)

func TestSuiteName(t *testing.T) {
	var suit Suite = 3
	var suit_name = suit.getSuiteName()

	var returned_name = "Clubs"

	if suit_name != returned_name {
		t.Errorf("There is some issue expected: %s, returned: %s", returned_name, suit_name)
	}
}

func TestFaceCardName(t *testing.T) {
	var face Face = 13
	var face_name = face.getFaceCard()

	var returned_name = "King"

	if face_name != returned_name {
		t.Errorf("There is some issue, expected: %s, returned: %s", returned_name, face_name)
	}
}

func TestCardDeck(t *testing.T) {
	var cards = New()

	fmt.Println(cards)
	if len(cards) != 52 {
		t.Errorf("There is some issue, expected length: 52, returned: %d", len(cards))
	}
}

func TestShuffleCards(t *testing.T) {
	var cards = New()
	var shuffledCards = Shuffle(cards)

	fmt.Println(shuffledCards)

	if shuffledCards[10] == cards[10] {
		t.Errorf("Cards are not shuffled")
	}
}

func TestStringerFunction(t *testing.T) {
	var card Card = Card{
		Suite: Diamonds,
		Face:  King,
	}

	var stringCard = card.String()

	if stringCard != "King of Diamonds" {
		t.Errorf("Expected %s Returned %s", "King of Diamonds", stringCard)
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Suite: Suite(Ace), Face: Face(Spades)}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades, received", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Suite: Suite(Ace), Face: Face(Spades)}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades, received", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suite == Joker {
			count++
		}
	}

	if count != 3 {
		t.Errorf("There should be 3 Jokers present on deck, the number of jokers present %d", count)
	}
}

func TestFilterFunc(t *testing.T) {
	filter := func(card Card) bool {
		return card.Face == Two || card.Face == Three
	}

	cards := New(FilterCards(filter))

	for _, c := range cards {
		if c.Face == Two || c.Face == Three {
			t.Errorf("All the Two and Three cards should be filtered")
		}
	}
}

func TestDecks(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Error("The number of cards should be 13*4*3, while the output is, ", len(cards))
	}
}
