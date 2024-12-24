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

func TestDay18Sample(t *testing.T) {
	t.Run("part 1 example 1", func(t *testing.T) {
		err := Day18(strings.NewReader(`
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`), 6, 6, 12)
		require.NoError(t, err)
	})
	t.Run("part 2 example 1", func(t *testing.T) {
		err := Day18(strings.NewReader(`
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`), 6, 6, -1)
		require.NoError(t, err)
	})
}

func TestDay18(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day18.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day18(f, 70, 70, 1024)
	require.NoError(t, err)

	f.Seek(0, io.SeekStart)

	err = Day18(f, 70, 70, -1)
	require.NoError(t, err)
}
