package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"aoc/util"
)

func Day10(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	var trailMap [][]int
	var rowN int
	var trailHeads [][2]int
	var mapWidth int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		if mapWidth == 0 {
			mapWidth = len(line)
		}

		var row []int
		for i, c := range line {
			height := util.MustInt(string(c))
			row = append(row, height)
			if height == 0 {
				trailHeads = append(trailHeads, [2]int{rowN, i})
			}
		}
		trailMap = append(trailMap, row)
		rowN++
	}

	movements := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	var total int
	var total2 int
	for _, trailHead := range trailHeads {
		toVisit := [][2]int{[2]int(trailHead[0:2])}
		uniquePeaks := map[[2]int]struct{}{}

		var rating int
		for len(toVisit) > 0 {
			pos := toVisit[len(toVisit)-1]
			toVisit = toVisit[:len(toVisit)-1]

			posHeight := trailMap[pos[0]][pos[1]]

			for _, movement := range movements {
				adjustedPos := [2]int{pos[0] + movement[0], pos[1] + movement[1]}
				if adjustedPos[0] < 0 || adjustedPos[0] >= len(trailMap) || adjustedPos[1] < 0 || adjustedPos[1] >= mapWidth {
					continue
				}

				if trailMap[adjustedPos[0]][adjustedPos[1]]-posHeight != 1 {
					continue
				} else if trailMap[adjustedPos[0]][adjustedPos[1]] == 9 {
					uniquePeaks[adjustedPos] = struct{}{}
					rating++
					continue
				}
				toVisit = append(toVisit, adjustedPos)
			}
		}
		total += len(uniquePeaks)
		total2 += rating
	}

	fmt.Println(total, total2)
	return nil
}
