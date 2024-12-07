package main

import (
	"slices"
	"testing"
)

func TestRanks(t *testing.T) {
	t.Run("five of a kind", func(t *testing.T) {
		got := rankHand(HandCards{'A', 'A', 'A', 'A', 'A'})
		expect := FiveOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("four of a kind", func(t *testing.T) {
		got := rankHand(HandCards{'A', 'A', '4', 'A', 'A'})
		expect := FourOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("full house", func(t *testing.T) {
		got := rankHand(HandCards{'A', 'A', '4', '4', 'A'})
		expect := FullHouse

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("Three of a kind", func(t *testing.T) {
		got := rankHand(HandCards{'A', '2', '4', 'A', 'A'})
		expect := ThreeOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("two pair", func(t *testing.T) {
		got := rankHand(HandCards{'A', 'A', '4', '4', '8'})
		expect := TwoPair

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("one pair", func(t *testing.T) {
		got := rankHand(HandCards{'2', 'A', '4', 'A', '9'})
		expect := OnePair

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("high card", func(t *testing.T) {
		got := rankHand(HandCards{'A', '6', '4', '@', '9'})
		expect := HighCard

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})
}

func TestHighCardWithJoekr(t *testing.T) {
	t.Run("one joker", func(t *testing.T) {
		got := rankHand(HandCards{'2', 'A', Joker, '4', '9'})
		expect := OnePair

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("two jokers", func(t *testing.T) {
		got := rankHand(HandCards{'2', Joker, Joker, '4', '9'})
		expect := ThreeOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("three jokers", func(t *testing.T) {
		got := rankHand(HandCards{'2', Joker, Joker, Joker, '9'})
		expect := FourOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("four jokers", func(t *testing.T) {
		got := rankHand(HandCards{Joker, Joker, Joker, Joker, '9'})
		expect := FiveOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("five jokers", func(t *testing.T) {
		got := rankHand(HandCards{Joker, Joker, Joker, Joker, Joker})
		expect := FiveOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

}
func TestPairWithJoker(t *testing.T) {
	t.Run("one joker", func(t *testing.T) {
		got := rankHand(HandCards{'2', 'A', Joker, 'A', '9'})
		expect := ThreeOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("two jokers", func(t *testing.T) {
		got := rankHand(HandCards{Joker, 'A', Joker, 'A', '9'})
		expect := FourOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})

	t.Run("three jokers", func(t *testing.T) {
		got := rankHand(HandCards{Joker, 'A', Joker, 'A', Joker})
		expect := FiveOfAKind

		if got != expect {
			t.Errorf("Got wrong hand, expected %d got %d", expect, got)
		}
	})
}

func TestSortHandCards(t *testing.T) {
	t.Run("Joker less than 2", func(t *testing.T) {
		expectedOrder := []byte{Joker, '2', '3', '4', '5', '6', '7', '8', '9', Ten, Jack, Queen, King, Ace}
		uut := []HandCards{
			{Joker, King, King, King, King},
			{'2', King, King, King, King},
		}
		for i := 0; i < len(expectedOrder)-1; i++ {
			uut[0][0] = expectedOrder[i]
			uut[1][0] = expectedOrder[i+1]
			slices.SortFunc(uut, sortHandCards)
			if uut[0][0] != expectedOrder[i+1] {
				t.Errorf("%d was bigger than %d", uut[0][0], expectedOrder[i+1])
			}
		}
	})
}
