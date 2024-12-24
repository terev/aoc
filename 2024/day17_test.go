package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay17Sample(t *testing.T) {
	t.Run("part 1 example 1", func(t *testing.T) {
		err := Day17(strings.NewReader(`
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`))
		require.NoError(t, err)
	})

	t.Run("part 2 example 1", func(t *testing.T) {
		err := Day17(strings.NewReader(`
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`))
		require.NoError(t, err)
	})
}

func TestDay17(t *testing.T) {
	err := Day17(strings.NewReader(`
Register A: 164278764924605
Register B: 0
Register C: 0

Program: 2,4,1,1,7,5,1,5,4,1,5,5,0,3,3,0`))
	require.NoError(t, err)
}
