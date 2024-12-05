package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay5_Sample(t *testing.T) {
	err := Day5(strings.NewReader(`
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,4`))
	require.NoError(t, err)
}

func TestDay5(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day5.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day5(f)
	require.NoError(t, err)
}
