package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"aoc/util"
)

type plantRegion struct {
	area       int
	boundaries map[[2]int][]int
}

var plotConnections = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var edgeIds = [][]int{{0, 2}, {1, 3}}

func findRegionAreaAndBoundary(start [2]int, farmMap []string) plantRegion {
	explored := map[[2]int]struct{}{}
	toExplore := [][2]int{start}

	plant := farmMap[start[0]][start[1]]

	boundaries := map[[2]int][]int{}

	for len(toExplore) > 0 {
		curPlot := toExplore[len(toExplore)-1]
		toExplore = toExplore[:len(toExplore)-1]
		if _, isExplored := explored[curPlot]; isExplored {
			continue
		}

		var edges []int
		for i, connection := range plotConnections {
			potentialConnection := [2]int{curPlot[0] + connection[0], curPlot[1] + connection[1]}
			if potentialConnection[0] < 0 || potentialConnection[0] >= len(farmMap) {
				edges = append(edges, i)
				continue
			}
			if potentialConnection[1] < 0 || potentialConnection[1] >= len(farmMap[potentialConnection[0]]) {
				edges = append(edges, i)
				continue
			}
			if farmMap[potentialConnection[0]][potentialConnection[1]] != plant {
				edges = append(edges, i)
				continue
			}
			toExplore = append(toExplore, potentialConnection)
		}

		if len(edges) > 0 {
			boundaries[curPlot] = edges
		}
		explored[curPlot] = struct{}{}
	}

	return plantRegion{
		boundaries: boundaries,
		area:       len(explored),
	}
}

// indicates if a point is within the polygon defined by a set of points along the edge.
// The given point is projected in each direction until a boundary is found.
func pointInRegion(p [2]int, boundaries map[[2]int][]int, farmMap []string) bool {
	if _, ok := boundaries[p]; ok {
		return true
	}

	plant := farmMap[p[0]][p[1]]

	boundariesToCheck := slices.Clone(plotConnections)
	boundariesFound := 0

	for i := 1; len(boundariesToCheck) > 0; i++ {
		var remainingBoundaries [][2]int
		for j := 0; j < len(boundariesToCheck); j++ {
			newP := [2]int{p[0] + boundariesToCheck[j][0]*i, p[1] + boundariesToCheck[j][1]*i}
			if _, ok := boundaries[newP]; ok {
				boundariesFound++
				continue
			}
			if newP[0] < 0 || newP[0] >= len(farmMap) ||
				newP[1] < 0 || newP[1] >= len(farmMap[0]) {
				return false
			}
			if farmMap[newP[0]][newP[1]] != plant {
				return false
			}
			remainingBoundaries = append(remainingBoundaries, boundariesToCheck[j])
		}
		boundariesToCheck = remainingBoundaries
	}
	return boundariesFound == 4
}

func countRegionSides(boundaryPoints map[[2]int][]int) int {
	// [0] = matching Ys
	// [1] = matching Xs
	matching := [2]map[int][]int{{}, {}}

	for p := range boundaryPoints {
		matching[0][p[0]] = append(matching[0][p[0]], p[1])
		matching[1][p[1]] = append(matching[1][p[1]], p[0])
	}

	var sides int

	for d, matches := range matching {
		// find runs of matching ys where the x is 1 apart
		for i, ps := range matches {
			slices.Sort(ps)

			thisPlot := [2]int{}
			thisPlot[d] = i

			lastP := -1
			for _, p := range ps {
				thisPlot[d^1] = p
				thisPlotEdges := util.IntersectSlices(boundaryPoints[thisPlot], edgeIds[d])

				if lastP == -1 {
					sides += len(thisPlotEdges)
					lastP = p
					continue
				}
				if p-lastP > 1 {
					sides += len(thisPlotEdges)
					lastP = p
					continue
				}

				lastCoord := thisPlot
				lastCoord[d^1] = lastP
				prevPlotEdges := util.IntersectSlices(boundaryPoints[lastCoord], edgeIds[d])
				sides += len(thisPlotEdges) - len(util.IntersectSlices(thisPlotEdges, prevPlotEdges))
				lastP = p
			}
		}
	}

	return sides
}

func Day12(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	var farmMap []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		farmMap = append(farmMap, line)
	}

	plantRegions := map[rune][]plantRegion{}

	for i := 0; i < len(farmMap); i++ {
	next:
		for j, plant := range farmMap[i] {
			plot := [2]int{i, j}

			for _, region := range plantRegions[plant] {
				if pointInRegion(plot, region.boundaries, farmMap) {
					continue next
				}
			}

			plantRegions[plant] = append(plantRegions[plant], findRegionAreaAndBoundary(plot, farmMap))
		}
	}

	var cost, cost2 int
	for plant, regions := range plantRegions {
		for _, region := range regions {
			var perimeter int
			for _, edges := range region.boundaries {
				perimeter += len(edges)
			}
			cost += region.area * perimeter
			sides := countRegionSides(region.boundaries)
			cost2 += region.area * sides

			fmt.Printf("Cost 1 %q: %d * %d = %d\n", plant, region.area, perimeter, region.area*perimeter)
			fmt.Printf("Cost 2 %q: %d * %d = %d\n", plant, region.area, sides, region.area*sides)
		}
	}

	fmt.Println(cost, cost2)
	return nil
}
