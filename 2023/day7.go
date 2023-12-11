package main

import (
	"aoc/util"
	"bufio"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"io"
	"strings"
)

func Day7(r io.Reader) (int, int, error) {
	type hand struct {
		cards string
		bid   int
	}
	var hands []hand
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		hands = append(hands, hand{
			cards: parts[0],
			bid:   util.MustInt(parts[1]),
		})
	}

	hands2 := slices.Clone(hands)

	slices.SortStableFunc(hands, func(a, b hand) bool {
		aType := handType(a.cards)
		bType := handType(b.cards)
		if aType == bType {
			return !compareHands(a.cards, b.cards)
		}
		if aType > bType {
			return true
		}
		return false
	})

	var winnings int
	for i := 0; i < len(hands); i++ {
		winnings += (1 + i) * hands[i].bid
	}

	slices.SortStableFunc(hands2, func(a, b hand) bool {
		aType := handTypeWithWildcard(a.cards)
		bType := handTypeWithWildcard(b.cards)
		if aType == bType {
			return !compareHandsWithWildcard(a.cards, b.cards)
		}
		if aType > bType {
			return true
		}
		return false
	})

	var winnings2 int
	for i := 0; i < len(hands2); i++ {
		winnings2 += (1 + i) * hands2[i].bid
	}

	return winnings, winnings2, nil
}

// returns true if a beats b
func compareHands(a, b string) bool {
	labelRanks := map[byte]int{
		'A': 1,
		'K': 2,
		'Q': 3,
		'J': 4,
		'T': 5,
		'9': 6,
		'8': 7,
		'7': 8,
		'6': 9,
		'5': 10,
		'4': 11,
		'3': 12,
		'2': 13,
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		return labelRanks[a[i]] < labelRanks[b[i]]
	}

	return false
}

// returns true if a beats b
func compareHandsWithWildcard(a, b string) bool {
	labelRanks := map[byte]int{
		'A': 1,
		'K': 2,
		'Q': 3,
		'T': 4,
		'9': 5,
		'8': 6,
		'7': 7,
		'6': 8,
		'5': 9,
		'4': 10,
		'3': 11,
		'2': 12,
		'J': 13,
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		return labelRanks[a[i]] < labelRanks[b[i]]
	}

	return false
}

func handType(hand string) int {
	kindCounts := make(map[rune]int)

	for _, r := range hand {
		kindCounts[r]++
	}

	counts := maps.Values(kindCounts)
	slices.SortFunc(counts, func(a, b int) bool {
		return b < a
	})

	if len(kindCounts) == 1 {
		return 1
	} else if len(counts) == 2 {
		if counts[0] == 4 {
			return 2
		}
		return 3
	} else if len(counts) == 3 {
		if counts[0] == 3 {
			return 4
		}
		return 5
	} else if len(counts) == 4 {
		return 6
	}
	return 7
}

func handTypeWithWildcard(hand string) int {
	kindCounts := make(map[rune]int)
	var wildCards int
	for _, r := range hand {
		if r == 'J' {
			wildCards++
			continue
		}
		kindCounts[r]++
	}

	if wildCards == 0 {
		return handType(hand)
	}

	counts := maps.Values(kindCounts)
	slices.SortFunc(counts, func(a, b int) bool {
		return b < a
	})

	if len(counts) == 1 || len(counts) == 0 {
		return 1
	} else if len(counts) == 2 { // AAAK, AAKK, AAK, AK
		if counts[0]+wildCards == 4 {
			return 2
		}
		return 3
	} else if len(counts) == 3 { // AAKT, AKT
		if counts[0]+wildCards == 3 {
			return 4
		}
		return 5
	} else if len(counts) == 4 { // AKT9
		return 6
	}

	panic("Nope")
}
