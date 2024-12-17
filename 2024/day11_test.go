package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unique"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay11Sample(t *testing.T) {
	err := Day11(strings.NewReader(`125 17`), 25)
	require.NoError(t, err)

	err = Day11(strings.NewReader(`125 17`), 75)
	require.NoError(t, err)
}

func TestNextStones(t *testing.T) {
	assert.ElementsMatch(t, util.UniqueSlice([]string{"1"}), nextStoneInscriptions(unique.Make("0")))
	assert.ElementsMatch(t, util.UniqueSlice([]string{"2", "0"}), nextStoneInscriptions(unique.Make("20")))

	assert.ElementsMatch(t, util.UniqueSlice([]string{"1", "7"}), nextStoneInscriptions(unique.Make("17")))
	assert.ElementsMatch(t, util.UniqueSlice([]string{"10", "7"}), nextStoneInscriptions(unique.Make("1007")))
	assert.ElementsMatch(t, util.UniqueSlice([]string{"253000"}), nextStoneInscriptions(unique.Make("125")))
}

func TestDay11(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day11.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day11(f, 25)
	require.NoError(t, err)

	_, err = f.Seek(io.SeekStart, 0)
	require.NoError(t, err)

	err = Day11(f, 75)
	require.NoError(t, err)
}
