package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"aoc/util"
)

func Day1(r io.ReadCloser) error {
	scanner := bufio.NewScanner(r)

	var lists [2][]int
	var ll int
	tally := make(map[int]int)

	for scanner.Scan() {
		l := scanner.Text()
		cols := strings.Fields(l)
		if len(cols) == 0 {
			continue
		} else if len(cols) != 2 {
			return fmt.Errorf("Invalid number of columns in input %d expected 2", len(cols))
		}

		for i, c := range cols {
			n, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				return err
			}
			ni := int(n)

			if i == 1 {
				tally[ni]++
			}

			lists[i] = append(lists[i], ni)
		}
		ll++
	}

	if err := r.Close(); err != nil {
		return err
	}

	for i := range lists {
		slices.Sort(lists[i])
	}

	var dist int
	var score int
	for i := 0; i < ll; i++ {
		dist += util.Abs(lists[1][i] - lists[0][i])
		if count, ok := tally[lists[0][i]]; ok {
			score += lists[0][i] * count
		}
	}
	fmt.Println("Distance:", dist)
	fmt.Println("Similarity:", score)

	return nil
}
