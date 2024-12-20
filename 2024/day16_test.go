package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay16Sample(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		minPathLength, minPathPositions, err := Day16(strings.NewReader(`
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`))
		require.NoError(t, err)
		assert.Equal(t, 7036, minPathLength)
		assert.Equal(t, 45, minPathPositions)
	})
	t.Run("example 2", func(t *testing.T) {
		minPathLength, minPathPositions, err := Day16(strings.NewReader(`
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`))
		require.NoError(t, err)
		assert.Equal(t, 11048, minPathLength)
		assert.Equal(t, 64, minPathPositions)
	})
}

func TestDay16(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day16.txt"))
	require.NoError(t, err)
	defer f.Close()

	minPathLength, minPathPositions, err := Day16(f)
	require.NoError(t, err)
	assert.Equal(t, 109496, minPathLength)
	assert.Equal(t, 551, minPathPositions)
}
