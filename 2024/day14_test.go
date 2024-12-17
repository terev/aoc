package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay14Sample(t *testing.T) {
	err := Day14(strings.NewReader(`
p=2,4 v=2,-3`), 11, 7, 5)
	require.NoError(t, err)

	err = Day14(strings.NewReader(`
	p=0,4 v=3,-3
	p=6,3 v=-1,-3
	p=10,3 v=-1,2
	p=2,0 v=2,-1
	p=0,0 v=1,3
	p=3,0 v=-2,-2
	p=7,6 v=-1,-3
	p=3,0 v=-1,-2
	p=9,3 v=2,3
	p=7,3 v=-1,2
	p=2,4 v=2,-3
	p=9,5 v=-3,-3`), 11, 7, 100)
	require.NoError(t, err)
}

func TestDay14(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day14.txt"))
	require.NoError(t, err)
	defer f.Close()

	// lower than 217654272
	err = Day14(f, 101, 103, 100)
	require.NoError(t, err)
}
