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

func TestDay2(t *testing.T) {
	s := `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	p1, p2, err := Day2(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 8, p1)
	assert.Equal(t, 2286, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day2.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, p2, err = Day2(f)
	require.NoError(t, err)
	fmt.Println(p1)
	fmt.Println(p2)
}
