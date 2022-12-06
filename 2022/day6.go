package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func findFirstWindowWithoutDups(haystack string, size int) int {
	lastSeen := [26]int{}
	windowStart := 0

	for i := 0; i < len(haystack); {
		id := haystack[i] - 97
		if ind := lastSeen[id]; ind > 0 {
			windowStart = ind
			i = ind
			lastSeen = [26]int{}
			continue
		} else if (i+1)-windowStart == size {
			return i + 1
		}
		lastSeen[id] = i + 1

		i++
	}

	return -1
}

func Day6() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day6.txt"))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		result := findFirstWindowWithoutDups(scanner.Text(), 4)
		fmt.Println(result)
		result2 := findFirstWindowWithoutDups(scanner.Text(), 14)
		fmt.Println(result2)
	}

	return nil
}
