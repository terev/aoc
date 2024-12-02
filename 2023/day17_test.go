package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDay17(t *testing.T) {
	s := `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

	p1, _, err := Day17(strings.NewReader(s))
	require.NoError(t, err)
	assert.Equal(t, 102, p1)
	//
	//f, err := os.Open(filepath.Join(util.Cwd(), "day17.txt"))
	//require.NoError(t, err)
	//defer f.Close()
	//
	//p1, _, err = Day17(f)
	//require.NoError(t, err)
	//fmt.Println(p1)
}
