package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay1(t *testing.T) {
	err := Day1(io.NopCloser(strings.NewReader(`
3   4
4   3
2   5
1   3
3   9
3   3
`)))
	require.NoError(t, err)

	f, err := os.Open(filepath.Join(util.Cwd(), "day1.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day1(f)
	require.NoError(t, err)
}
