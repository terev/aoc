package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
)

const (
	sensorPattern  = `Sensor at x=(-?\d+),\s*y=(-?\d+): closest beacon is at x=(-?\d+),\s*y=(-?\d+)`
	checkYPosition = 2000000
	maxBeaconX     = 4000000
	maxBeaconY     = 4000000
)

var sensorRe *regexp.Regexp

func init() {
	sensorRe = regexp.MustCompile(sensorPattern)
}

func manhattan(a, b point) int {
	return util.Abs(a.x-b.x) + util.Abs(a.y-b.y)
}

type sensor struct {
	pos       point
	beaconPos point
	distance  int
	vertices  [4]point
}

type point struct {
	x int
	y int
}

func addVector(p point, vector [2]int) point {
	return point{
		x: p.x + vector[0],
		y: p.y + vector[1],
	}
}

func Day15() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day15.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	var sensors []sensor

	xBounds := [2]int{math.MaxInt64, math.MinInt64}

	lineScanner := bufio.NewScanner(f)
	for lineScanner.Scan() {
		matches := sensorRe.FindStringSubmatch(lineScanner.Text())
		if len(matches) == 0 {
			continue
		}

		sensorPos := point{util.MustInt(matches[1]), util.MustInt(matches[2])}
		beaconPos := point{util.MustInt(matches[3]), util.MustInt(matches[4])}

		beaconDistance := manhattan(sensorPos, beaconPos)

		sensors = append(sensors, sensor{
			pos:       sensorPos,
			beaconPos: beaconPos,
			distance:  beaconDistance,
			vertices: [4]point{
				// left
				{sensorPos.x - beaconDistance, sensorPos.y},
				// top
				{sensorPos.x, sensorPos.y - beaconDistance},
				// right
				{sensorPos.x + beaconDistance, sensorPos.y},
				// bottom
				{sensorPos.x, sensorPos.y + beaconDistance},
			},
		})

		for _, v := range sensors[len(sensors)-1].vertices {
			if v.x < xBounds[0] {
				xBounds[0] = v.x
			}

			if v.x > xBounds[1] {
				xBounds[1] = v.x
			}
		}
	}

	var notBeaconPositions = map[[2]int]struct{}{}

	for x := xBounds[0]; x <= xBounds[1]; x++ {
		for _, scanner := range sensors {
			if (x != scanner.beaconPos.x || checkYPosition != scanner.beaconPos.y) &&
				manhattan(scanner.pos, point{x, checkYPosition}) <= scanner.distance {
				notBeaconPositions[[2]int{x, checkYPosition}] = struct{}{}
			}
		}
	}

	fmt.Println(len(notBeaconPositions))

	zigZagVectors := [4][2]int{
		{1, -1},
		{1, 1},
		{-1, 1},
		{-1, -1},
	}

	initVectors := [4][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	candiates := map[[2]int]struct{}{}

	for _, scanner := range sensors {
		for i, vertex := range scanner.vertices {
			endVertex := (i + 1) % (len(scanner.vertices))
			cur := addVector(vertex, initVectors[i])
			end := addVector(scanner.vertices[endVertex], initVectors[endVertex])

			for cur.x != end.x || cur.y != end.y {
				if cur.x >= 0 && cur.x <= maxBeaconX && cur.y >= 0 && cur.y <= maxBeaconY {
					candiates[[2]int{cur.x, cur.y}] = struct{}{}
				}
				cur = addVector(cur, zigZagVectors[i])
			}
		}
	}

	for _, scanner := range sensors {
		for p := range candiates {
			if manhattan(scanner.pos, point{p[0], p[1]}) <= scanner.distance {
				delete(candiates, p)
			}
		}
	}

	for p := range candiates {
		fmt.Println(p[0]*4000000 + p[1])
	}

	return nil
}
