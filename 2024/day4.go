package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"aoc/util"
)

var xmas = "XMAS"

func Day4(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	var grid [][]int
	var row []int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if row == nil {
			row = make([]int, len(line))
		}
		for i, c := range line {
			row[i] = strings.IndexRune(xmas, c)
		}

		grid = append(grid, slices.Clone(row))
	}

	height, width := len(grid), len(row)

	var tot int
	for i := range height {
		tot += countXMAS(grid[i])
	}

	for i := range width {
		var col []int
		for j := range height {
			col = append(col, grid[j][i])
		}
		tot += countXMAS(col)
	}

	for i := 0; i <= width+height-2; i++ {
		var diag []int

		for j := 0; j <= i; j++ {
			k := i - j

			if k < height && j < width {
				diag = append(diag, grid[k][j])
			}
		}

		tot += countXMAS(diag)
	}

	for i := 0; i <= width+height-2; i++ {
		var diag []int

		for j := width - 1; j >= 0; j-- {
			k := i - j
			if k >= 0 && k < height && j < width {
				diag = append(diag, grid[k][j])
			}
		}

		tot += countXMAS(diag)
	}

	fmt.Println(tot)

	fmt.Println(countXPattern(grid))

	return nil
}

func countXMAS(sector []int) int {
	var asc *bool
	var prev *int

	var count int

	for _, char := range sector {
		if prev == nil {
			if char == 0 || char == 3 {
				prev = &char
			}
			continue
		}

		diff := *prev - char
		sign := diff < 0

		if util.Abs(diff) != 1 || (asc != nil && sign != *asc) {
			asc = nil
			if char != 0 && char != 3 {
				prev = nil
			} else {
				prev = &char
			}
			continue
		}

		if char == 0 || char == 3 {
			count++
			asc = nil
		} else if asc == nil {
			asc = &sign
		}

		prev = &char
	}
	return count
}

func countXPattern(grid [][]int) int {
	var tot int
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] != 2 {
				continue
			}
			// Found a possible X center point, check for MS.

			left := []int{grid[i-1][j-1], grid[i+1][j+1]}
			slices.Sort(left)
			if !slices.Equal([]int{1, 3}, left) {
				continue
			}
			right := []int{grid[i+1][j-1], grid[i-1][j+1]}
			slices.Sort(right)

			if !slices.Equal([]int{1, 3}, right) {
				continue
			}
			tot++
		}
	}
	return tot
}
