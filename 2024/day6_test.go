package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay6Sample(t *testing.T) {
	err := Day6(strings.NewReader(`
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`))
	require.NoError(t, err)
}

func TestDay6(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day6.txt"))
	require.NoError(t, err)
	defer f.Close()
	err = Day6(f)
	require.NoError(t, err)
}
