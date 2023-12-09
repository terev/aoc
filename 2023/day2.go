package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func Day2(r io.Reader) (int, int, error) {

	var sum, sum2 int
	var gameID int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, fmt.Sprintf("Game %d: ", gameID+1))
		if isGamePossible(line) {
			sum += gameID + 1
		}

		mins := minForPossibleGame(line)
		sum2 += mins[0] * mins[1] * mins[2]

		gameID++
	}

	return sum, sum2, nil
}

func isGamePossible(game string) bool {
	sets := strings.Split(game, ";")
	for _, set := range sets {
		results := strings.Split(set, ",")
		for _, result := range results {
			parts := strings.Split(strings.TrimSpace(result), " ")
			n, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}
			switch parts[1] {
			case "red":
				if n > maxRed {
					return false
				}
			case "green":
				if n > maxGreen {
					return false
				}
			case "blue":
				if n > maxBlue {
					return false
				}
			}
		}
	}

	return true
}

func minForPossibleGame(game string) [3]int {
	var mins [3]int
	sets := strings.Split(game, ";")
	for _, set := range sets {
		results := strings.Split(set, ",")
		for _, result := range results {
			parts := strings.Split(strings.TrimSpace(result), " ")
			n, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}
			switch parts[1] {
			case "red":
				mins[0] = max(mins[0], int(n))
			case "green":
				mins[1] = max(mins[1], int(n))
			case "blue":
				mins[2] = max(mins[2], int(n))
			}
		}
	}

	return mins
}
