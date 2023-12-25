package main

import (
	"aoc/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDay16(t *testing.T) {
	s := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	p1, p2, err := Day16(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 46, p1)
	assert.Equal(t, 51, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day16.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, p2, err = Day16(f)
	assert.Equal(t, 7392, p1)
	assert.Equal(t, 7665, p2)
}
