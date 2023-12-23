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

func TestDay13(t *testing.T) {
	s := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

`

	p1, p2, err := Day13(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 405, p1)
	assert.Equal(t, 400, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day13.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, p2, err = Day13(f)
	require.NoError(t, err)
	assert.Equal(t, 34993, p1)
	assert.Equal(t, 29341, p2)
}
