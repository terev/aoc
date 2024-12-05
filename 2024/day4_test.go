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

func TestCountXMAS(t *testing.T) {
	assert.Equal(t, 1, countXMAS([]int{0, 1, 2, 3}))
	assert.Equal(t, 1, countXMAS([]int{3, 2, 1, 0}))
	assert.Equal(t, 3, countXMAS([]int{0, 1, 2, 3, 3, 3, 0, 1, 2, 3, 0, 3, 0, 1, 2, 3, 0}))
}

func TestDay4Sample(t *testing.T) {
	err := Day4(strings.NewReader(`
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`))
	require.NoError(t, err)
}

func TestDay4(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day4.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day4(f)
	require.NoError(t, err)
}
