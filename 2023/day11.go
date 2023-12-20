package main

import (
	"aoc/util"
	"bufio"
	"golang.org/x/exp/slices"
	"io"
)

var expansionRate = 10

func Day11(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	var image []string
	for scanner.Scan() {
		image = append(image, scanner.Text())
	}

	var galaxies [][2]int

	columnsNotEmpty := make([]bool, len(image[0]))
	var emptyRows []int
	for i := 0; i < len(image); i++ {
		rowEmpty := true
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == '#' {
				rowEmpty = false
				columnsNotEmpty[j] = true
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, i)
		}
	}

	var emptyColumns []int
	for i, notEmptyColumn := range columnsNotEmpty {
		if !notEmptyColumn {
			emptyColumns = append(emptyColumns, i)
		}
	}

	galaxies2 := slices.Clone(galaxies)

	for offset, i := range emptyRows {
		for j := range galaxies {
			if galaxies[j][0] > i+offset {
				galaxies[j][0]++
			}
			if galaxies2[j][0] > i+offset*expansionRate-offset {
				galaxies2[j][0] += expansionRate - 1
			}
		}
	}

	for offset, i := range emptyColumns {
		for j := range galaxies {
			if galaxies[j][1] > i+offset {
				galaxies[j][1]++
			}

			if galaxies2[j][1] > i+offset*expansionRate-offset {
				galaxies2[j][1] += expansionRate - 1
			}
		}
	}

	var sum, sum2 int
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := util.Abs(galaxies[i][1]-galaxies[j][1]) + util.Abs(galaxies[i][0]-galaxies[j][0])
			distance2 := util.Abs(galaxies2[i][1]-galaxies2[j][1]) + util.Abs(galaxies2[i][0]-galaxies2[j][0])

			sum += distance
			sum2 += distance2
		}
	}

	return sum, sum2, nil
}
