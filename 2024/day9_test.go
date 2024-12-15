package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay9Sample(t *testing.T) {
	err := Day9(strings.NewReader(`2333133121414131402`))
	require.NoError(t, err)
}

func TestDay9(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day9.txt"))
	require.NoError(t, err)
	defer f.Close()
	err = Day9(f)
	require.NoError(t, err)
}
