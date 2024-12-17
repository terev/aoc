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

func TestDay13Sample(t *testing.T) {
	err := Day13(strings.NewReader(`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`))

	require.NoError(t, err)
}

func TestParseCoord(t *testing.T) {
	assert.Equal(t, [2]int{17, 86}, parseCoord("X+17, Y+86"))
	assert.Equal(t, [2]int{18641, 10279}, parseCoord("X=18641, Y=10279"))
}

func TestDay13(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day13.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day13(f)
	require.NoError(t, err)
}
