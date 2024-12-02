package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"github.com/chen3feng/stl4go"
	"io"
)

func Day17(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)
	var heatmap [][]int
	for scanner.Scan() {
		var row []int
		for _, c := range scanner.Bytes() {
			row = append(row, int(c-48))
		}
		heatmap = append(heatmap, row)
	}

	minHeatLoss := traverseCity(heatmap)

	return minHeatLoss, 0, nil
}

type cityTraversalNode struct {
	position           [2]int
	incurredHeatLoss   int
	straightMovements  int
	enteredFrom        int
	traveledDirections []int
}

func manhattan(p1, p2 [2]int) int {
	return util.Abs(p1[0]-p2[0]) + util.Abs(p1[1]-p2[1])
}

func traverseCity(heatmap [][]int) int {
	target := [2]int{len(heatmap) - 1, len(heatmap[0]) - 1}
	toVisit := stl4go.NewPriorityQueueFunc[cityTraversalNode](func(a, b cityTraversalNode) bool {
		return a.incurredHeatLoss+manhattan(a.position, target) < b.incurredHeatLoss+manhattan(b.position, target)
	})
	toVisit.Push(cityTraversalNode{position: [2]int{}, incurredHeatLoss: 0, straightMovements: 0, enteredFrom: -1})

	directionalMovementVectors := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	visited := make(map[[2]int]struct{})
	minHeatLoss := -1

	for toVisit.Len() > 0 {
		current := toVisit.Pop()
		//visited[current.position] = struct{}{}

		//fmt.Println(current.position)

		for i := 0; i < 4; i++ {
			if i == current.enteredFrom {
				continue
			}

			dirIsStraightLine := current.enteredFrom != -1 && i == (current.enteredFrom+2)%4
			if current.straightMovements == 3 && dirIsStraightLine {
				continue
			}

			neighborPos := [2]int{
				current.position[0] + directionalMovementVectors[i][0],
				current.position[1] + directionalMovementVectors[i][1],
			}

			if neighborPos[0] < 0 || neighborPos[1] < 0 || neighborPos[0] >= len(heatmap) || neighborPos[1] >= len(heatmap[0]) {
				continue
			}

			if neighborPos[0] == target[0] && neighborPos[1] == target[1] {
				for _, dir := range current.traveledDirections {
					switch dir {
					case 0:
						fmt.Println("⌃")
					case 1:
						fmt.Println(">")
					case 2:
						fmt.Println("˅")
					case 3:
						fmt.Println("<")
					}
				}
				fmt.Println("PATH END")
				if minHeatLoss == -1 {
					minHeatLoss = current.incurredHeatLoss + int(heatmap[neighborPos[0]][neighborPos[1]])
				} else {
					minHeatLoss = min(minHeatLoss, current.incurredHeatLoss+int(heatmap[neighborPos[0]][neighborPos[1]]))
				}
				continue
			}

			if _, isVisited := visited[neighborPos]; isVisited {
				continue
			}
			var straightMovements = 1
			if dirIsStraightLine {
				straightMovements = current.straightMovements + 1
			}

			visited[neighborPos] = struct{}{}
			toVisit.Push(cityTraversalNode{
				position:           neighborPos,
				incurredHeatLoss:   current.incurredHeatLoss + heatmap[neighborPos[0]][neighborPos[1]],
				straightMovements:  straightMovements,
				enteredFrom:        (i + 2) % 4,
				traveledDirections: append(current.traveledDirections, i),
			})
		}
	}

	return minHeatLoss
}
