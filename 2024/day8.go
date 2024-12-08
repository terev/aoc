package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Day8(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	uniqueAntinodes1 := map[[2]int]struct{}{}
	uniqueAntinodes2 := map[[2]int]struct{}{}
	antennae := map[byte][][2]int{}
	var width, row int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if width == 0 {
			width = len(line)
		}
		for i, c := range line {
			if c == '.' {
				continue
			}
			antennae[byte(c)] = append(antennae[byte(c)], [2]int{i, row})
		}
		row++
	}
	height := row

	for _, positions := range antennae {
		if len(positions) < 2 {
			continue
		}
		for i := 0; i < len(positions); i++ {
			a := positions[i]
			for j := i + 1; j < len(positions); j++ {
				distX := positions[j][0] - a[0]
				distY := positions[j][1] - a[1]

				for rep, pos := range getAntinodesInRange(a, distX, distY, width, height) {
					if rep == 1 {
						uniqueAntinodes1[pos] = struct{}{}
					}
					uniqueAntinodes2[pos] = struct{}{}
				}

				for rep, pos := range getAntinodesInRange(positions[j], -distX, -distY, width, height) {
					if rep == 1 {
						uniqueAntinodes1[pos] = struct{}{}
					}
					uniqueAntinodes2[pos] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(uniqueAntinodes1))
	fmt.Println(len(uniqueAntinodes2))

	return nil
}

func getAntinodesInRange(antanaPos [2]int, freqX, freqY, rangeWidth, rangeHeight int) [][2]int {
	var antinodePositions [][2]int

	for mult := 0; ; mult++ {
		antinodePos := [2]int{antanaPos[0] - (freqX * mult), antanaPos[1] - (freqY * mult)}

		if antinodePos[0] < 0 || antinodePos[0] >= rangeWidth ||
			antinodePos[1] < 0 || antinodePos[1] >= rangeHeight {
			return antinodePositions
		}

		antinodePositions = append(antinodePositions, antinodePos)
	}
	return antinodePositions
}
