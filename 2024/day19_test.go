package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay19Sample(t *testing.T) {
	t.Run("part 1 example 1", func(t *testing.T) {
		err := Day19(strings.NewReader(`
r, wr, b, g, bwu, rb, gb, br

rrbgbr
brwrr
bggr
gbbr
ubwu
bwurrg
brgr
bbrgwb`))
		require.NoError(t, err)
	})
}

func TestDay19(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day19.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day19(f)
	require.NoError(t, err)
}
