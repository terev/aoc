package main

import (
	"aoc/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandType(t *testing.T) {
	typeTests := []struct {
		hand         string
		expectedType int
	}{
		{
			"AAAAA",
			1,
		},
		{
			"AA8AA",
			2,
		},
		{
			"23332",
			3,
		},
		{
			"TTT98",
			4,
		},
		{
			"23432",
			5,
		},
		{
			"A23A4",
			6,
		},
		{
			"23456",
			7,
		},
	}

	for _, test := range typeTests {
		assert.Equal(t, test.expectedType, handType(test.hand))
	}
}

func TestHandTypeWithWildCards(t *testing.T) {
	typeTests := []struct {
		hand         string
		expectedType int
	}{
		{
			"AAAAA",
			1,
		},
		{
			"AAAAJ",
			1,
		},
		{
			"AA8AA",
			2,
		},
		{
			"AA8JA",
			2,
		},
		{
			"23332",
			3,
		},
		{
			"23J32",
			3,
		},
		{
			"TTT98",
			4,
		},
		{
			"JTT98",
			4,
		},
		{
			"23432",
			5,
		},
		{
			"J3432",
			4,
		},
		{
			"A23A4",
			6,
		},
		{
			"J23A4",
			6,
		},
		{
			"23456",
			7,
		},
		{
			"2345J",
			6,
		},
		{
			"234JJ",
			4,
		},
		{
			"2222J",
			1,
		},
		{
			"T55J5",
			2,
		},
	}

	for _, test := range typeTests {
		t.Run(test.hand, func(t *testing.T) {
			assert.Equal(t, test.expectedType, handTypeWithWildcard(test.hand))
		})
	}
}

func TestDay7(t *testing.T) {
	s := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	p1, p2, err := Day7(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 6440, p1)
	assert.Equal(t, 5905, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day7.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, p2, err = Day7(f)
	require.NoError(t, err)
	fmt.Println(p1)
	fmt.Println(p2)
}
