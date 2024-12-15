package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay10Sample(t *testing.T) {
	err := Day10(strings.NewReader(`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`))

	require.NoError(t, err)
}

func TestDay10(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day10.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day10(f)
	require.NoError(t, err)
}
