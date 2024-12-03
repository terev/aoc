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

func TestDay3(t *testing.T) {
	err := Day3P1(strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`))
	require.NoError(t, err)

	err = Day3P2(strings.NewReader(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`))
	require.NoError(t, err)

	f, err := os.Open(filepath.Join(util.Cwd(), "day3.txt"))
	require.NoError(t, err)
	defer f.Close()
	err = Day3P1(f)
	require.NoError(t, err)

	_, err = f.Seek(0, io.SeekStart)
	require.NoError(t, err)

	err = Day3P2(f)
	require.NoError(t, err)
}
