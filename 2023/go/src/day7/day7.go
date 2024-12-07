package main

import (
	"advent2023/pkg/numerics"
	"fmt"
	"path/filepath"
	"slices"
	"time"

	"advent2023/pkg/file"
)

type HandRank int
type HandCards [5]byte

// We are going to transform the face cards into over values so that they are ordered
const Ten = '@'
const Jack = 'J'
const Queen = 'Q'
const King = 'V'
const Ace = 'Z'
const Joker = '0'

const (
	HighCard     HandRank = iota
	OnePair      HandRank = iota
	TwoPair      HandRank = iota
	ThreeOfAKind HandRank = iota
	FullHouse    HandRank = iota
	FourOfAKind  HandRank = iota
	FiveOfAKind  HandRank = iota
)

type CamelHand struct {
	Hand HandCards
	Bid  int
}

func sortHandCards(a HandCards, b HandCards) int {
	for i := 0; i < 5; i++ {
		if a[i] != b[i] {
			return int(b[i]) - int(a[i])
		}
	}
	return 0
}

func sortCamelHand(a CamelHand, b CamelHand) int {
	return sortHandCards(a.Hand, b.Hand)
}

func rankHand(hand HandCards) HandRank {
	cardCounts := make(map[byte]int)

	// It'll be fine
	maxCount := 0
	jokerCount := 0
	for i := 0; i < 5; i++ {
		if hand[i] != Joker {
			cardCounts[hand[i]] += 1
			if cardCounts[hand[i]] > maxCount {
				maxCount = cardCounts[hand[i]]
			}
		} else {
			jokerCount++
		}
	}

	maxCount += jokerCount

	switch maxCount {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if len(cardCounts) == 2 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	case 2:
		if len(cardCounts) == 3 {
			return TwoPair
		} else {
			return OnePair
		}
	default:
		return HighCard
	}
}

func part1(lines []string) float64 {
	hands := make(map[HandRank][]CamelHand)

	for _, line := range lines {
		current := CamelHand{}
		for i := 0; i < 5; i++ {
			c := line[i]
			// It will make it easier to sort later if we make sure that the value of the card is ordered.  2 - 9
			// will already be in the correct order then J Q are also ok.  So we just need to move K and A to make sure
			// the value of those cards is greater than Q.
			if c == 'K' {
				c = King
			} else if c == 'A' {
				c = Ace
			} else if c == 'T' {
				c = '@'
			}
			current.Hand[i] = c
		}
		current.Bid, _ = numerics.GetNumeric(line, 5)

		rank := rankHand(current.Hand)
		hands[rank] = append(hands[rank], current)
	}

	ranks := [7]HandRank{FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPair, OnePair, HighCard}
	position := len(lines)
	result := 0.0
	for _, rank := range ranks {
		handGroup := hands[rank]
		slices.SortFunc(handGroup, sortCamelHand)
		for _, hand := range handGroup {
			result += float64(position) * float64(hand.Bid)
			position--
		}
	}

	if position != 0 {
		fmt.Printf("Position is %d", position)
	}

	return result
}

func part2(lines []string) float64 {
	hands := make(map[HandRank][]CamelHand)

	for _, line := range lines {
		current := CamelHand{}
		for i := 0; i < 5; i++ {
			c := line[i]
			switch c {
			case 'T':
				c = Ten
			case 'J':
				c = Joker
			case 'Q':
				c = Queen
			case 'K':
				c = King
			case 'A':
				c = Ace
			}
			current.Hand[i] = c
		}
		current.Bid, _ = numerics.GetNumeric(line, 5)

		rank := rankHand(current.Hand)
		hands[rank] = append(hands[rank], current)
	}

	ranks := [7]HandRank{FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPair, OnePair, HighCard}
	position := len(lines)
	result := 0.0
	for _, rank := range ranks {
		handGroup := hands[rank]
		slices.SortFunc(handGroup, sortCamelHand)
		for _, hand := range handGroup {
			result += float64(position) * float64(hand.Bid)
			position--
		}
	}

	if position != 0 {
		fmt.Printf("Position is %d", position)
	}

	return result
}

func main() {
	abs, _ := filepath.Abs("input")
	output, _ := file.ReadInput(abs)

	start := time.Now()
	result := part1(output)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 done with %f. It took %s\n", result, elapsed)

	start = time.Now()
	result = part2(output)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 done with %f. It took %s\n", result, elapsed)
}
