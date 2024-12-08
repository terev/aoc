package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"aoc/util"
)

func Day7(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	var p1, p2 int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		testValRaw, numsRaw, cut := strings.Cut(line, ":")
		if !cut {
			panic("Colon not found")
		}

		testVal := util.MustInt(testValRaw)
		nums := strings.Fields(strings.TrimSpace(numsRaw))

		var totals = []int{util.MustInt(nums[0])}
		var totals2 = []int{util.MustInt(nums[0])}
		for i := 1; i < len(nums); i++ {
			newNum := util.MustInt(nums[i])
			var newTotals []int
			for _, t := range totals {
				newTotals = append(newTotals, t+newNum)
				newTotals = append(newTotals, t*newNum)
			}
			totals = newTotals

			var newTotals2 []int
			for _, t := range totals2 {
				newTotals2 = append(newTotals2, t+newNum)
				newTotals2 = append(newTotals2, t*newNum)
				newTotals2 = append(newTotals2, concatB10Nums(t, newNum))
			}
			totals2 = newTotals2
		}

		if slices.Contains(totals, testVal) {
			p1 += testVal
		}

		if slices.Contains(totals2, testVal) {
			p2 += testVal
		}
	}

	fmt.Println(p1, p2)
	return nil
}

func concatB10Nums(a, b int) int {
	concat := a
	divs := 10
	for b/divs > 0 {
		divs *= 10
	}
	return concat*divs + b
}
