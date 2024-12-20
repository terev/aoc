package main

import (
	"bufio"
	"container/heap"
	"io"
	"maps"
	"math"
	"slices"

	"aoc/util"
)

func Day16(in io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(in)
	tm, err := readTileMap(scanner)
	if err != nil {
		return 0, 0, err
	}

	start := slices.Collect(maps.Keys(tm.TypeLocations['S']))[0]
	end := slices.Collect(maps.Keys(tm.TypeLocations['E']))[0]

	minPathLength, minPathPositions := minimalPath(tm, start, end)

	return minPathLength, minPathPositions, nil
}

const (
	movePenalty   = 1
	rotatePenalty = 1000
)

func minimalPath(tm tileMap, start, end [2]int) (minPathLength int, minPathTiles int) {
	vertices := slices.Collect(maps.Keys(tm.TypeLocations[emptyTile]))
	vertices = append(vertices, start, end)
	pq := util.PriorityQueue[[3]int]{}

	// [row, column, lookDir] -> [minDist]
	distances := map[[3]int]int{}
	paths := map[[3]int][][3]int{}

	startVec := [3]int{start[0], start[1], 1}
	pq.PushValue(startVec, 0)
	distances[startVec] = 0

	for _, pos := range vertices {
		for j := range 4 {
			posVec := [3]int{pos[0], pos[1], j}
			if posVec == startVec {
				continue
			} else {
				pq.PushValue(posVec, math.MaxInt)
				distances[posVec] = math.MaxInt
			}
		}
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		cur, prevScore := pq.PopValue()

		for i, neighborDir := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			dirDiff := util.Abs(i - cur[2])
			if dirDiff == 3 {
				dirDiff = 1
			}

			neighborPos := [2]int{cur[0] + neighborDir[0], cur[1] + neighborDir[1]}
			neighborTile, ok := tm.Lookup[neighborPos]
			if !ok {
				continue
			}
			score := prevScore

			var actionVec [3]int
			if dirDiff == 0 {
				// Neighbor straight ahead.
				if neighborTile == wallTile {
					// Can't explore through walls.
					continue
				}
				// Empty space in looked direction, can potentially explore forward.
				score += movePenalty
				actionVec = [3]int{neighborPos[0], neighborPos[1], i}
			} else {
				// Must rotate from current looked dir.
				score += rotatePenalty * dirDiff
				actionVec = [3]int{cur[0], cur[1], i}
			}

			if score < distances[actionVec] {
				distances[actionVec] = score
				paths[actionVec] = [][3]int{cur}
				pq.UpdatePriority(actionVec, score)
			} else if score == distances[actionVec] {
				paths[actionVec] = append(paths[actionVec], cur)
			}
		}
	}

	var minPathVec [3]int
	minScore := math.MaxInt
	for vec, distance := range distances {
		if vec[0] == end[0] && vec[1] == end[1] && distance < minScore {
			minScore = distance
			minPathVec = vec
		}
	}

	tilesOnLeastPaths := map[[2]int]struct{}{}
	tilesOnLeastPaths[start] = struct{}{}
	tilesOnLeastPaths[end] = struct{}{}
	prevPaths := paths[minPathVec]
	for len(prevPaths) > 0 {
		p := prevPaths[len(prevPaths)-1]
		prevPaths = prevPaths[:len(prevPaths)-1]
		if ps, hasPaths := paths[p]; hasPaths {
			prevPaths = append(prevPaths, ps...)
		}
		tilesOnLeastPaths[[2]int{p[0], p[1]}] = struct{}{}
	}

	return minScore, len(tilesOnLeastPaths)
}
