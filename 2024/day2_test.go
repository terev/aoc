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

func TestDay2(t *testing.T) {
	err := Day2(io.NopCloser(strings.NewReader(`
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`)))
	require.NoError(t, err)

	var f io.ReadCloser
	f, err = os.Open(filepath.Join(util.Cwd(), "day2.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day2(f)
	require.NoError(t, err)
}
