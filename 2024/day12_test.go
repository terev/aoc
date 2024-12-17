package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"aoc/util"
)

func TestDay12Sample(t *testing.T) {
	err := Day12(strings.NewReader(`
AAAA
BBCD
BBCC
EEEC`))
	require.NoError(t, err)

	err = Day12(strings.NewReader(`
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`))
	require.NoError(t, err)
	err = Day12(strings.NewReader(`
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`))
	require.NoError(t, err)

	err = Day12(strings.NewReader(`
OOOOOHHHHHHHOHHHHQQQRQRDDDDBBBBBBBBBBBBYYYYYYYYYYYYYYYXXXXXXXXXXXXXXQQQQQTTTTTM
OOOOOOHHOOOOOHHHQQQQQQQDDBBBBBBBBBBBBIBVYYYYYYYYYYYYYYJXXXXXXXXXXXXQQQQQQQMMTTM
OOOOOOOOOOOOOHXQQQQQQQDDDDBBBBBBBBBBBBBBBQYYYYYYYYYYYJJXXXXXXXXXXXXXQQQPPQMTTTM
OOOOOOOOOOOOOHQQQQQQQZZZDDBBBBBBBBBBBBBBQQYYYEYYYYYYYYYXXXXXXXXXXXXXNPPNPMMMMMM
OOWOOOOOOOOOOQQQQQQQZZZZZDBBBBBBBBBBBSBQQQYYYEYYYYYYYFFXXXXXXXXXMMXNNNNNNNNNMNN
WWWOOOOXOXXOOQQQQQQZZZZZZBBYBBBBBBBBBBBQQEEEEEMYYYIVVIFXXXXXXMMMMMMNNNNNNNNNNNN
WWWWWWOXXXXQQQQQQQQQZZZZZBBBBBBBBBBBBBBBEEEEEEEYYYIIIIIIXLMMXMMMMMMNNNNNNNNNNNN`))
	require.NoError(t, err)
}

func TestDay12(t *testing.T) {
	f, err := os.Open(filepath.Join(util.Cwd(), "day12.txt"))
	require.NoError(t, err)
	defer f.Close()

	err = Day12(f)
	require.NoError(t, err)
}

func TestCountRegionSides(t *testing.T) {
	type regionSideCheck struct {
		name          string
		regionStart   [2]int
		expectedSides int
	}
	testCases := []struct {
		mapData      string
		regionChecks []regionSideCheck
	}{
		{
			mapData: `
AAAA
BBCD
BBCC
EEEC`,
			regionChecks: []regionSideCheck{
				{
					name:          "A",
					regionStart:   [2]int{0, 0},
					expectedSides: 4,
				},
				{
					name:          "B",
					regionStart:   [2]int{1, 0},
					expectedSides: 4,
				},
				{
					name:          "C",
					regionStart:   [2]int{3, 3},
					expectedSides: 8,
				},
			},
		},
		{
			mapData: `
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			regionChecks: []regionSideCheck{
				{
					name:          "E",
					regionStart:   [2]int{0, 0},
					expectedSides: 12,
				},
			},
		},
	}

	for i, testCase := range testCases {
		mapData := strings.Split(strings.TrimSpace(testCase.mapData), "\n")
		for _, check := range testCase.regionChecks {
			t.Run(fmt.Sprintf("%d_%s", i, check.name), func(t *testing.T) {
				region := findRegionAreaAndBoundary(check.regionStart, mapData)
				assert.Equal(t, check.expectedSides, countRegionSides(region.boundaries))
			})
		}
	}
}
