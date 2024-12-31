package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"maps"
	"slices"

	"aoc/util"
)

func Day20(in io.Reader, maxCheatDuration int) error {
	scanner := bufio.NewScanner(in)
	tm, err := readTileMap(scanner)
	if err != nil {
		return err
	}

	start := slices.Collect(maps.Keys(tm.TypeLocations['S']))[0]
	end := slices.Collect(maps.Keys(tm.TypeLocations['E']))[0]
	minPath := raceLengths(tm, start, end)

	cheatingTimeSaves := findCheatsPaths(tm, end, minPath, maxCheatDuration)

	var metThreshold int
	for timeSaved, count := range cheatingTimeSaves {
		if timeSaved > 0 {
			fmt.Println(count, "route(s) saved", timeSaved)
			if timeSaved >= 100 {
				metThreshold += count
			}
		}
	}
	fmt.Println(metThreshold, "met threshold")
	return nil
}

func raceLengths(tm tileMap, start, end [2]int) map[[2]int]int {
	distances := map[[2]int]int{}

	distance := 0
	cur := start

	for cur != end {
		distances[cur] = distance

		for _, neighborDir := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			neighborPos := [2]int{cur[0] + neighborDir[0], cur[1] + neighborDir[1]}
			if _, ok := distances[neighborPos]; ok {
				continue
			}

			neighborTile, ok := tm.Lookup[neighborPos]
			if !ok {
				continue
			}

			if neighborTile != wallTile {
				cur = neighborPos
				break
			}
		}
		distance++
	}

	distances[end] = distance

	return distances
}

type pathKey struct {
	pos           [2]int
	cheatStartPos [2]int
}

func findCheatsPaths(tm tileMap, endPos [2]int, minPath map[[2]int]int, maxCheatDuration int) map[int]int {
	pq := util.PriorityQueue[pathKey]{}
	distances := map[pathKey]int{}

	for pos, prevPosDist := range minPath {
		if pos == endPos {
			continue
		}

		pk := pathKey{pos: pos, cheatStartPos: pos}
		distances[pk] = prevPosDist
		pq.PushValue(pk, prevPosDist)
	}
	heap.Init(&pq)

	timeSaves := map[int]int{}
	maxStartEnds := map[pathKey]int{}
	for len(pq) > 0 {
		cur, prevScore := pq.PopValue()
		if distances[cur] != prevScore {
			break
		}

		newMinDistance := prevScore + 1

		for _, neighborVec := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			neighborPos := [2]int{cur.pos[0] + neighborVec[0], cur.pos[1] + neighborVec[1]}
			if neighborPos[0] < 0 || neighborPos[1] < 0 ||
				neighborPos[0] >= tm.Height || neighborPos[1] >= tm.Width {
				continue
			}

			neighborTile := tm.Lookup[neighborPos]

			next := cur
			next.pos = neighborPos

			if prevDistance, ok := distances[next]; !ok || newMinDistance < prevDistance {
				cheatDuration := util.ManhattanDistance(next.cheatStartPos, next.pos)
				if neighborTile != wallTile && cheatDuration <= maxCheatDuration {
					timeSave := minPath[next.pos] - newMinDistance
					if prevSaved, hasPrev := maxStartEnds[next]; !hasPrev {
						maxStartEnds[next] = timeSave
						timeSaves[timeSave]++
					} else if timeSave > prevSaved {
						maxStartEnds[next] = timeSave
						timeSaves[timeSave]++
						timeSaves[prevSaved]--
					}
				}
				distances[next] = newMinDistance

				if cheatDuration < maxCheatDuration {
					pq.PushValue(next, newMinDistance)
				}
			}
		}
	}

	return timeSaves
}
