package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck")
	}
}
func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade} // Creating expression so it will be properly evaluated
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Recieved:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Recieved:", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// frist call to shuffleRand.Perm(52) should be
	// [40 35...]
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]

	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected first card to be %s recieved %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected second card to be  %s recieved %s", second, cards[1])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jocers, recived:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 ranks * 4 suits * 3 deks
	if len(cards) != 13*4*3 {
		t.Errorf("Exoected %d cards, got %d cards.", 13*4*3, len(cards))
	}
}
