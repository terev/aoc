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

func TestDay9(t *testing.T) {
	s := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	p1, p2, err := Day9(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 114, p1)
	assert.Equal(t, 2, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day9.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, p2, err = Day9(f)
	require.NoError(t, err)
	fmt.Println(p1)
	fmt.Println(p2)
}
