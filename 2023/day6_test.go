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

func TestDay6(t *testing.T) {
	s := `Time:      7  15   30
Distance:  9  40  200`
	p1, p2, err := Day6(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 288, p1)
	assert.Equal(t, 71503, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day6.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, p2, err = Day6(f)
	require.NoError(t, err)

	fmt.Println(p1)
	fmt.Println(p2)
}
