package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay7Sample(t *testing.T) {
	err := Day7(strings.NewReader(`
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`))
	require.NoError(t, err)
}

func TestDay7(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day7.txt"))
	require.NoError(t, err)
	defer f.Close()
	err = Day7(f)
	require.NoError(t, err)
}

func TestConcatB10Nums(t *testing.T) {
	assert.Equal(t, 101, concatB10Nums(10, 1))
	assert.Equal(t, 10120, concatB10Nums(10, 120))
}
