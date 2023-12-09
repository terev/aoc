package main

import (
	"aoc/util"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestDay1(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day1.txt"))
	require.NoError(t, err)

	p1, p2, err := Day1(f)
	require.NoError(t, err)

	fmt.Println(p1, p2)
}
