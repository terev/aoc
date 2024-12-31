package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay20Sample(t *testing.T) {
	example1 := `
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...# 
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
	t.Run("part 1 example 1", func(t *testing.T) {
		err := Day20(strings.NewReader(example1), 2)
		require.NoError(t, err)
	})
	t.Run("part 2 example 1", func(t *testing.T) {
		err := Day20(strings.NewReader(example1), 20)
		require.NoError(t, err)
	})
}

func TestDay20P1(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day20.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day20(f, 2)
	require.NoError(t, err)
}

func TestDay20P2(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day20.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day20(f, 20)
	require.NoError(t, err)
}
