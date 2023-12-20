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

func TestDay11(t *testing.T) {
	s := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	p1, p2, err := Day11(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 374, p1)
	assert.Equal(t, 1030, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day11.txt"))
	require.NoError(t, err)
	defer f.Close()

	expansionRate = 1_000_000
	p1, p2, err = Day11(f)
	require.NoError(t, err)
	fmt.Println(p1)
	fmt.Println(p2)
}
