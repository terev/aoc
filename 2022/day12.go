package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func squareValue(c int) int {
	if c == 'E' {
		return 'z'
	} else if c == 'S' {
		return 'a'
	}
	return c
}

func minimalPath(heightMap [][]int, start [2]int) (int, bool) {
	var toVisit = [][3]int{{0, start[0], start[1]}}
	visited := map[[2]int]struct{}{[2]int{start[0], start[1]}: {}}

	mapWidth := len(heightMap[0])
	for len(toVisit) > 0 {
		cur := toVisit[0]
		if len(toVisit) > 1 {
			toVisit = toVisit[1:]
		} else {
			toVisit = [][3]int{}
		}

		curSquare := squareValue(heightMap[cur[1]][cur[2]])
		for _, neighbor := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			ni := cur[1] + neighbor[0]
			nj := cur[2] + neighbor[1]
			if ni < 0 || ni > len(heightMap)-1 ||
				nj < 0 || nj > mapWidth-1 {
				continue
			}

			if _, alreadySeen := visited[[2]int{ni, nj}]; alreadySeen {
				continue
			}

			if squareValue(heightMap[ni][nj])-curSquare > 1 {
				continue
			}

			if heightMap[ni][nj] == 'E' {
				return cur[0] + 1, true
			}

			toVisit = append(toVisit, [3]int{cur[0] + 1, ni, nj})
			// mark neighbour visited to avoid being discovered from another node
			visited[[2]int{ni, nj}] = struct{}{}
		}
	}

	return -1, false
}

func Day12() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day12.txt"))
	if err != nil {
		return err
	}
	defer f.Close()
	var heightmap [][]int
	scanner := bufio.NewScanner(f)
	var start [2]int
	i := 0
	for scanner.Scan() {
		heightmap = append(heightmap, []int{})
		for j, c := range scanner.Text() {
			if c == 'S' {
				start = [2]int{i, j}
			}

			heightmap[i] = append(heightmap[i], int(c))
		}
		i++
	}

	minFromStart, _ := minimalPath(heightmap, start)
	fmt.Println(minFromStart)

	min := minFromStart
	for i := range heightmap {
		for j := range heightmap[i] {
			if heightmap[i][j] == 'a' {
				minStepsFromHere, reachable := minimalPath(heightmap, [2]int{i, j})
				if reachable && minStepsFromHere < min {
					min = minStepsFromHere
				}
			}
		}
	}

	fmt.Println(min)

	return nil
}
