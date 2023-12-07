package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/mymmrac/x"
)

var cardsScore = map[string]int64{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var cardsScoreJ = map[string]int64{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

type Hand struct {
	Cards       []string
	Bid         int64
	Combination int
}

func part1(input string) any {
	lines := strings.Split(input, "\n")

	var hands []Hand
	for _, line := range lines {
		var cards string
		var bid int64
		_, err := fmt.Sscanf(line, "%s %d", &cards, &bid)
		x.Assert(err == nil, err)

		hands = append(hands, Hand{
			Cards: strings.Split(cards, ""),
			Bid:   bid,
		})
	}

	for i, hand := range hands {
		cards := map[string]int64{}
		for _, card := range hand.Cards {
			cards[card]++
		}

		var combinations []int64
		for _, count := range cards {
			combinations = append(combinations, count)
		}
		slices.Sort(combinations)
		slices.Reverse(combinations)

		switch len(cards) {
		case 1: // Five of a kind
			hand.Combination = 7
		case 2: // Four of a kind or Full house
			if combinations[0] == 4 { // Four of a kind
				hand.Combination = 6
			} else { // Full house
				hand.Combination = 5
			}
		case 3: // Three of a kind or Two pair
			if combinations[0] == 3 { // Three of a kind
				hand.Combination = 4
			} else { // Two pair
				hand.Combination = 3
			}
		case 4: // One pair
			hand.Combination = 2
		case 5: // High card
			hand.Combination = 1
		}
		hands[i] = hand
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		h := cmp.Compare(b.Combination, a.Combination)
		if h == 0 {
			for i := range b.Cards {
				c := cmp.Compare(cardsScore[b.Cards[i]], cardsScore[a.Cards[i]])
				if c != 0 {
					return c
				}
			}
			panic("unreachable")
		}
		return h
	})

	var answer int64
	for i, hand := range hands {
		answer += hand.Bid * int64(len(hands)-i)
	}

	return answer
}

type Comb struct {
	Card  string
	Count int64
}

func part2(input string) any {
	lines := strings.Split(input, "\n")

	var hands []Hand
	for _, line := range lines {
		var cards string
		var bid int64
		_, err := fmt.Sscanf(line, "%s %d", &cards, &bid)
		x.Assert(err == nil, err)

		hands = append(hands, Hand{
			Cards: strings.Split(cards, ""),
			Bid:   bid,
		})
	}

	for i, hand := range hands {
		cards := map[string]int64{}
		for _, card := range hand.Cards {
			cards[card]++
		}

		var combinations []Comb
		for card, count := range cards {
			combinations = append(combinations, Comb{
				Card:  card,
				Count: count,
			})
		}
		slices.SortFunc(combinations, func(a, b Comb) int {
			c := cmp.Compare(b.Count, a.Count)
			if c != 0 {
				return c
			}
			return cmp.Compare(cardsScoreJ[b.Card], cardsScoreJ[a.Card])
		})

		switch len(cards) {
		case 1: // Five of a kind
			hand.Combination = 7
		case 2: // Four of a kind or Full house
			if combinations[0].Count == 4 { // Four of a kind
				hand.Combination = 6
				if countJ(hand.Cards) == 1 || countJ(hand.Cards) == 4 {
					hand.Combination = 7
				}
			} else { // Full house
				hand.Combination = 5
				if countJ(hand.Cards) == 2 || countJ(hand.Cards) == 3 {
					hand.Combination = 7
				}
			}
		case 3: // Three of a kind or Two pair
			if combinations[0].Count == 3 { // Three of a kind
				hand.Combination = 4
				switch countJ(hand.Cards) {
				case 1:
					hand.Combination = 6
				case 3:
					hand.Combination = 6
				}
			} else { // Two pair
				hand.Combination = 3
				switch countJ(hand.Cards) {
				case 1:
					hand.Combination = 5
				case 2:
					hand.Combination = 6
				}
			}
		case 4: // One pair
			hand.Combination = 2
			switch countJ(hand.Cards) {
			case 1:
				hand.Combination = 4
			case 2:
				hand.Combination = 4
			}
		case 5: // High card
			hand.Combination = 1
			if countJ(hand.Cards) == 1 {
				hand.Combination = 2
			}
		}
		hands[i] = hand
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		h := cmp.Compare(b.Combination, a.Combination)
		if h == 0 {
			for i := range b.Cards {
				c := cmp.Compare(cardsScoreJ[b.Cards[i]], cardsScoreJ[a.Cards[i]])
				if c != 0 {
					return c
				}
			}
			panic("unreachable")
		}
		return h
	})

	var answer int64
	for i, hand := range hands {
		answer += hand.Bid * int64(len(hands)-i)
	}

	return answer
}

func countJ(cards []string) int {
	c := 0
	for _, card := range cards {
		if card == "J" {
			c++
		}
	}
	return c
}
