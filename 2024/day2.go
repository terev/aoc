package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"aoc/util"
)

func Day2(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	var safe int
	var safeWithBad int

read:
	for scanner.Scan() {
		levelsRaw := strings.Fields(scanner.Text())
		if len(levelsRaw) == 0 {
			continue
		}

		initialError, valid := validateReport(levelsRaw)
		if valid {
			safe++
			continue
		}

		for _, inc := range []int{-1, 0, 1} {
			ri := initialError + inc
			if ri < 0 || ri > len(levelsRaw) {
				continue
			}

			s := slices.Delete(slices.Clone(levelsRaw), ri, ri+1)
			if _, valid := validateReport(s); valid {
				safeWithBad++
				continue read
			}
		}
	}

	fmt.Println("Safe:", safe)
	fmt.Println("SafeWithBad:", safe+safeWithBad)

	return nil
}

func validateReport(report []string) (int, bool) {
	var asc *bool
	var prev *int

	for i := 0; i < len(report); i++ {
		level := util.MustInt(report[i])

		if prev == nil {
			prev = &level
			continue
		}

		diff := *prev - level
		sign := diff < 0

		if diff == 0 || util.Abs(diff) > 3 {
			return i - 1, false
		}
		if asc != nil && sign != *asc {
			return i - 1, false
		}

		*prev = level
		if asc == nil {
			asc = &sign
		}
	}

	return 0, true
}
