package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"aoc/util"
)

var buttonRegex = regexp.MustCompile(`X[+=](\d+)[^Y]*Y[+=](\d+)`)

func Day13(in io.Reader) error {
	var minCost, minCost2 int

	var prize, buttonA, buttonB [2]int

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		prefix, leftover, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		switch prefix {
		case "Button A":
			buttonA = parseCoord(leftover)
		case "Button B":
			buttonB = parseCoord(leftover)
		case "Prize":
			prize = parseCoord(leftover)
			cost := minCostToWinPrize(prize, buttonA, buttonB)
			if cost != -1 {
				minCost += cost
			}

			prize = [2]int{
				prize[0] + 10000000000000,
				prize[1] + 10000000000000,
			}
			cost = minCostToWinPrize(prize, buttonA, buttonB)
			if cost != -1 {
				minCost2 += cost
			}
		}
	}

	fmt.Println(minCost, minCost2)

	return nil
}

func parseCoord(coordRaw string) [2]int {
	parts := buttonRegex.FindStringSubmatch(coordRaw)
	return [2]int{util.MustInt(parts[1]), util.MustInt(parts[2])}
}

func minCostToWinPrize(target [2]int, buttonA [2]int, buttonB [2]int) int {
	minCost := -1

	// solve:
	// axi+bxj - tx = 0
	// ayi+byj - ty = 0

	j := (buttonB[1]*buttonA[0] - buttonB[0]*buttonA[1]) * buttonA[0]
	tgt := util.Abs(buttonA[1]*target[0]-target[1]*buttonA[0]) * buttonA[0]
	j = tgt / j
	i := util.Abs(j*buttonB[0]-target[0]) / buttonA[0]

	if i*buttonA[0]+j*buttonB[0] == target[0] &&
		i*buttonA[1]+j*buttonB[1] == target[1] {
		minCost = i*3 + j
	}

	k := (buttonB[0]*buttonA[1] - buttonB[1]*buttonA[0]) * buttonA[1]
	tgt2 := util.Abs(buttonB[1]*target[0]-target[1]*buttonB[0]) * buttonA[1]
	k = tgt2 / k
	l := util.Abs(k*buttonA[0]-target[0]) / buttonB[0]

	if k*buttonA[0]+l*buttonB[0] == target[0] &&
		k*buttonA[1]+l*buttonB[1] == target[1] {
		cost := k*3 + l
		if minCost == -1 || minCost > cost {
			return cost
		}
	}

	return minCost
}
