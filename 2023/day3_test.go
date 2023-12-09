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

func TestDay3(t *testing.T) {
	s := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	p1, p2, err := Day3(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 4361, p1)
	assert.Equal(t, 467835, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day3.txt"))
	require.NoError(t, err)
	p1, p2, err = Day3(f)
	fmt.Println(p1)
	fmt.Println(p2)
}
