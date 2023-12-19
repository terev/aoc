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

func TestDay8(t *testing.T) {
	s := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	p1, err := Day8(strings.NewReader(s))
	require.NoError(t, err)

	assert.Equal(t, 2, p1)

	s = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	p1, err = Day8(strings.NewReader(s))
	require.NoError(t, err)

	assert.Equal(t, 6, p1)

	s = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	p2, err := Day82(strings.NewReader(s))
	require.NoError(t, err)

	assert.Equal(t, 6, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day8.txt"))
	require.NoError(t, err)

	p1, err = Day8(f)
	require.NoError(t, err)

	fmt.Println(p1)
	f.Close()

	f, err = os.Open(filepath.Join(util.Cwd(), "day8.txt"))
	require.NoError(t, err)

	p2, err = Day82(f)
	require.NoError(t, err)

	fmt.Println(p2)
	f.Close()
}
