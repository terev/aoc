package main

import (
	"aoc/util"
	"bufio"
	"golang.org/x/exp/slices"
	"io"
	"strings"
)

func Day9(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)
	var sum, sum2 int
	for scanner.Scan() {
		sequenceData := strings.Fields(scanner.Text())
		var originalSequence []int

		for _, s := range sequenceData {
			originalSequence = append(originalSequence, util.MustInt(s))
		}

		curSeq := slices.Clone(originalSequence)

		var sequenceZeros = false
		var seqi int
		for !sequenceZeros {
			var newSequence []int
			sequenceZeros = true
			for i := len(curSeq) - 1; i > 0; i-- {
				diff := curSeq[i] - curSeq[i-1]
				newSequence = slices.Insert(newSequence, 0, diff)
				if diff != 0 {
					sequenceZeros = false
				}
			}

			sum += curSeq[len(curSeq)-1]
			if seqi%2 == 0 {
				sum2 += curSeq[0]
			} else {
				sum2 -= curSeq[0]
			}
			seqi++
			curSeq = newSequence
		}
	}
	return sum, sum2, nil
}
