package main

import (
	"aoc/util"
	"bufio"
	"golang.org/x/exp/slices"
	"io"
)

func Day10(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)
	var grid []string
	var start [2]int
	connections := map[rune][]int{
		'S': {0, 1, 2, 3},
		'L': {0, 1},
		'J': {0, 3},
		'7': {2, 3},
		'F': {1, 2},
		'|': {0, 2},
		'-': {1, 3},
	}

	for scanner.Scan() {
		line := scanner.Text()
		for j, r := range line {
			if r == 'S' {
				start = [2]int{len(grid), j}
			}
		}
		grid = append(grid, line)
	}

	type node struct {
		pos   [2]int
		steps int
		path  [][2]int
	}

	var toVisit = []node{
		{
			pos:   start,
			steps: 0,
			path:  [][2]int{},
		},
	}

	visited := make(map[[2]int]struct{})

	var maxLoop = 0
	var maxPath [][2]int
	for len(toVisit) > 0 {
		cur := toVisit[0]
		toVisit = slices.Delete(toVisit, 0, 1)
		visited[cur.pos] = struct{}{}

		for _, connectionPoint := range connections[rune(grid[cur.pos[0]][cur.pos[1]])] {
			var neighborPos = cur.pos
			switch connectionPoint {
			case 0:
				neighborPos[0] -= 1
			case 1:
				neighborPos[1] += 1
			case 2:
				neighborPos[0] += 1
			case 3:
				neighborPos[1] -= 1
			}
			if neighborPos[0] < 0 || neighborPos[0] > len(grid) || neighborPos[1] < 0 || neighborPos[1] > len(grid[0]) {
				continue
			}
			neighbor := grid[neighborPos[0]][neighborPos[1]]
			if neighbor == 'S' && cur.steps > 1 {
				if cur.steps+1 > maxLoop {
					maxLoop = cur.steps + 1
					maxPath = append(cur.path, cur.pos)
				}
				continue
			}

			if _, isVisited := visited[neighborPos]; isVisited {
				continue
			}

			if neighborConnections, ok := connections[rune(neighbor)]; ok && slices.Contains(neighborConnections, (connectionPoint+2)%4) {
				toVisit = slices.Insert(toVisit, 0, node{
					pos:   neighborPos,
					steps: cur.steps + 1,
					path:  append(cur.path, cur.pos),
				})
			}
		}
	}

	var corners [][2]int = [][2]int{maxPath[0]}

	for i := 0; i < len(maxPath); i++ {
		if slices.Contains([]byte{'7', 'L', 'J', 'F'}, grid[maxPath[i][0]][maxPath[i][1]]) {
			corners = append(corners, maxPath[i])
		}
	}

	corners = append(corners, maxPath[0])

	// Shoelace formula
	var sum int
	for i := 0; i < len(corners); i++ {
		sum += (len(grid) - corners[i][0]) * (corners[circularIndex(len(corners), i-1)][1] - corners[circularIndex(len(corners), i+1)][1])
	}

	return maxLoop / 2, (util.Abs(sum)-maxLoop)/2 + 1, nil
}

func circularIndex(size int, i int) int {
	if i < 0 {
		return size + i
	}
	if i >= size {
		return i % size
	}

	return i
}
