package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"maps"
	"math"
	"os"
	"slices"
	"strings"

	"aoc/util"
)

func Day18(in io.Reader, memoryWidth, memoryHeight, simMemory int) error {
	tm := tileMap{
		EmptyTile: emptyTile,
		Width:     memoryWidth + 1,
		Height:    memoryHeight + 1,
	}
	tm.Clear()

	scanner := bufio.NewScanner(in)
	i := 0
	for scanner.Scan() {
		if simMemory != -1 && i > simMemory {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		x, y, found := strings.Cut(line, ",")
		if !found {
			continue
		}

		pos := [2]int{util.MustInt(y), util.MustInt(x)}

		tm.SetTile(pos, wallTile)

		if simMemory == -1 {
			minp, _ := minimalPathThroughMemory(tm, [2]int{0, 0}, [2]int{memoryWidth, memoryHeight})
			if minp == -1 {
				fmt.Printf("%s,%s\n", x, y)
				return nil
			}
		}

		i++
	}

	tm.WriteTo(os.Stdout)

	fmt.Println(minimalPathThroughMemory(tm, [2]int{0, 0}, [2]int{memoryWidth, memoryHeight}))

	return nil
}

func minimalPathThroughMemory(tm tileMap, start, end [2]int) (minPathLength int, minPathTiles int) {
	vertices := slices.Collect(maps.Keys(tm.TypeLocations[emptyTile]))
	vertices = append(vertices, start, end)
	pq := util.PriorityQueue[[2]int]{}

	// [row, column] -> [minDist]
	distances := map[[2]int]int{}
	paths := map[[2]int][][2]int{}

	pq.PushValue(start, 0)
	distances[start] = 0

	for _, pos := range vertices {
		if pos == start {
			continue
		} else {
			pq.PushValue(pos, math.MaxInt)
			distances[pos] = math.MaxInt
		}
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		cur, prevScore := pq.PopValue()
		if prevScore == math.MaxInt {
			break
		}

		for _, neighborDir := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			neighborPos := [2]int{cur[0] + neighborDir[0], cur[1] + neighborDir[1]}
			if neighborPos[0] < 0 || neighborPos[0] > tm.Height || neighborPos[1] < 0 || neighborPos[1] > tm.Width {
				continue
			}
			neighborTile, ok := tm.Lookup[neighborPos]
			if !ok || neighborTile != tm.EmptyTile {
				continue
			}
			score := prevScore + movePenalty

			if score < distances[neighborPos] {
				distances[neighborPos] = score
				paths[neighborPos] = [][2]int{cur}
				pq.UpdatePriority(neighborPos, score)
			} else if score == distances[neighborPos] {
				paths[neighborPos] = append(paths[neighborPos], cur)
			}
		}
	}

	if _, ok := paths[end]; !ok {
		return -1, 0
	}

	return distances[end], 0
}
