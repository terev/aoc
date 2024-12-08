package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Day6(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	obstacles := map[[2]int]struct{}{}

	guard := [3]int{}
	var row, width int
	for scanner.Scan() {
		gridRow := strings.TrimSpace(scanner.Text())
		if len(gridRow) == 0 {
			continue
		}
		if width == 0 {
			width = len(gridRow)
		}

		for col, tile := range gridRow {
			switch tile {
			case '^':
				guard = [3]int{col, row, 0}
			case '>':
				guard = [3]int{col, row, 1}
			case 'v':
				guard = [3]int{col, row, 2}
			case '<':
				guard = [3]int{col, row, 3}
			case '#':
				obstacles[[2]int{col, row}] = struct{}{}
			}
		}
		row++
	}

	width, height := width-1, row-1
	start := [2]int{guard[0], guard[1]}

	path, loop := moveGuardUntilLoopOrEscape(guard, obstacles, width, height)
	if loop {
		panic("Initial guard path loops")
	}

	uniquePositions := map[[2]int]struct{}{}
	loopObstacles := map[[2]int]struct{}{}
	for _, ghost := range path {
		uniquePositions[[2]int(ghost[0:2])] = struct{}{}

		potentialObstaclePos := nextInFacedDirection(ghost)
		if potentialObstaclePos == start || isObstacle(obstacles, potentialObstaclePos) {
			continue
		}

		obstacles[potentialObstaclePos] = struct{}{}
		if _, loop := moveGuardUntilLoopOrEscape(guard, obstacles, width, height); loop {
			loopObstacles[potentialObstaclePos] = struct{}{}
		}
		delete(obstacles, potentialObstaclePos)
	}

	fmt.Println(len(uniquePositions), len(loopObstacles))
	return nil
}

func moveGuardUntilLoopOrEscape(guard [3]int, obstacles map[[2]int]struct{}, width, height int) ([][3]int, bool) {
	visited := map[[3]int]struct{}{}
	var path [][3]int

	for !canEscape(guard, width, height) {
		path = append(path, guard)
		newPos := nextInFacedDirection(guard)
		if isObstacle(obstacles, newPos) {
			guard[2] = (guard[2] + 1) % 4
			visited[guard] = struct{}{}
			continue
		}

		if _, ok := visited[[3]int{newPos[0], newPos[1], guard[2]}]; ok {
			return path, true
		}

		guard[0], guard[1] = newPos[0], newPos[1]
		visited[guard] = struct{}{}
	}

	path = append(path, guard)

	return path, false
}

func isObstacle(obstacles map[[2]int]struct{}, p [2]int) bool {
	_, ya := obstacles[p]
	return ya
}

func canEscape(guard [3]int, width, height int) bool {
	return (guard[2] == 0 && guard[1] == 0) ||
		(guard[2] == 2 && guard[1] == height) ||
		(guard[2] == 3 && guard[0] == 0) ||
		(guard[2] == 1 && guard[0] == width)
}

func nextInFacedDirection(guard [3]int) [2]int {
	switch guard[2] {
	case 0, 2:
		delta := guard[2] - 1
		return [2]int{guard[0], guard[1] + delta}
	case 1, 3:
		delta := guard[2] - 2
		return [2]int{guard[0] - delta, guard[1]}
	}
	panic(fmt.Errorf("not a dir: %d", guard[2]))
}
