package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func itemPriority(item byte) int {
	if item >= 97 {
		return int(item - 96)
	} else {
		return int(item - 38)
	}
}

func Day3() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day3.txt"))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	var priorityTotal int
	var sackId int = 1
	groupTally := map[byte][3]bool{}
	var groupPriorityTotal int

	for scanner.Scan() {
		t := scanner.Text()

		firstCompartment := map[byte]struct{}{}
		mid := len(t) / 2

		added := false

		for i, c := range []byte(t) {
			if tally, ok := groupTally[c]; ok {
				tally[(sackId-1)%3] = true
				groupTally[c] = tally
			} else {
				var tally = [3]bool{}
				tally[(sackId-1)%3] = true
				groupTally[c] = tally
			}

			if i < mid {
				firstCompartment[c] = struct{}{}
			} else if !added {
				if _, shared := firstCompartment[c]; shared {
					priorityTotal += itemPriority(c)
					added = true
				}
			}
		}

		if sackId%3 == 0 {
			for c, sacks := range groupTally {
				if sacks[0] && sacks[1] && sacks[2] {
					groupPriorityTotal += itemPriority(c)
					break
				}
			}

			groupTally = map[byte][3]bool{}
		}

		sackId++
	}

	fmt.Println(priorityTotal)
	fmt.Println(groupPriorityTotal)

	return nil
}
