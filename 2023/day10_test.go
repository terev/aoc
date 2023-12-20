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

	s = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

	_, p2, err := Day10(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 10, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day10.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, p2, err = Day10(f)
	require.NoError(t, err)
	fmt.Println(p1)
	fmt.Println(p2)
}
