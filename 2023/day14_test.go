package main

import (
	"aoc/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDay14(t *testing.T) {
	s := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	p1, _, err := Day14(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 136, p1)

	f, err := os.Open(filepath.Join(util.Cwd(), "day14.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, _, err = Day14(f)
	require.NoError(t, err)
	assert.Equal(t, 110677, p1)
}
