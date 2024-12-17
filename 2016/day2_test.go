package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay2Sample(t *testing.T) {
	err := Day2(strings.NewReader(`
ULL
RRDDD
LURDL
UUUUD`))
	require.NoError(t, err)
}

func TestDay2(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day2.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day2(f)
	require.NoError(t, err)
}
