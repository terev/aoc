package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Day4(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	copiesWon := make(map[int]int)

	var sum, sum2 int
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ":")

		cardID, err := strconv.ParseInt(strings.Fields(lineParts[0])[1], 10, 64)
		if err != nil {
			return 0, 0, err
		}

		sets := strings.Split(lineParts[1], "|")
		winners := make(map[int]struct{})

		for _, ns := range strings.Fields(strings.TrimSpace(sets[0])) {
			n, err := strconv.ParseInt(ns, 10, 64)
			if err != nil {
				return 0, 0, err
			}
			winners[int(n)] = struct{}{}
		}

		var numberWon int
		var points int
		for _, ns := range strings.Fields(strings.TrimSpace(sets[1])) {
			n, err := strconv.ParseInt(ns, 10, 64)
			if err != nil {
				return 0, 0, err
			}

			if _, ok := winners[int(n)]; ok {
				numberWon++
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
		sum2 += 1

		intCardID := int(cardID)

		for i := intCardID + 1; i <= intCardID+numberWon; i++ {
			copiesWon[i] += 1 + copiesWon[intCardID]
		}
	}

	for _, copies := range copiesWon {
		sum2 += copies
	}

	return sum, sum2, nil
}
