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

func TestDay15(t *testing.T) {
	assert.Equal(t, 52, hash(`HASH`))

	s := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
	p1, p2, err := Day15(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 1320, p1)
	assert.Equal(t, 145, p2)

	f, err := os.Open(filepath.Join(util.Cwd(), "day15.txt"))
	require.NoError(t, err)
	defer f.Close()

	p1, p2, err = Day15(f)
	require.NoError(t, err)
	fmt.Println(p1)
	fmt.Println(p2)
}
