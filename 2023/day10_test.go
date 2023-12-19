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

func TestDay10(t *testing.T) {
	s := `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`
	p1, _, err := Day10(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 4, p1)

	s = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

	_, p2, err := Day10(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 4, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day10.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, _, err = Day10(f)
	require.NoError(t, err)
	fmt.Println(p1)
}
