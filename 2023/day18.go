package main

import (
	"aoc/util"
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Day18(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)
	var corners [][2]int
	var corners2 [][2]int
	var cursor [2]int
	var cursor2 [2]int

	perimeter := 0
	perimeter2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		instruction := strings.Fields(line)
		steps := util.MustInt(instruction[1])

		hexData := strings.Trim(instruction[2], "(#)")
		distance, err := strconv.ParseInt(hexData[0:5], 16, 64)
		if err != nil {
			return 0, 0, err
		}
		direction := hexData[5] - 48

		corners = append(corners, cursor)
		corners2 = append(corners2, cursor2)

		prevPoint := cursor
		prevPoint2 := cursor2

		switch instruction[0] {
		case "L":
			cursor[1] -= steps
		case "R":
			cursor[1] += steps
		case "U":
			cursor[0] -= steps
		case "D":
			cursor[0] += steps
		}

		switch direction {
		case 0:
			cursor2[1] += int(distance)
		case 1:
			cursor2[0] += int(distance)
		case 2:
			cursor2[1] -= int(distance)
		case 3:
			cursor2[0] -= int(distance)
		}

		perimeter += util.Abs(prevPoint[0]-cursor[0]) + util.Abs(prevPoint[1]-cursor[1])
		perimeter2 += util.Abs(prevPoint2[0]-cursor2[0]) + util.Abs(prevPoint2[1]-cursor2[1])
	}
	corners = append(corners, [2]int{})
	corners2 = append(corners2, [2]int{})

	area := (util.AreaOfPolygon(corners)+perimeter)/2 + 1
	area2 := (util.AreaOfPolygon(corners2)+perimeter2)/2 + 1
	return area, area2, nil
}
