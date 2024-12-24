package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"strings"
)

const (
	robotTile              = '@'
	wallTile               = '#'
	boxTile                = 'O'
	emptyTile              = '.'
	doubleWideBoxLeftTile  = '['
	doubleWideBoxRightTile = ']'
)

var (
	instructionMovements = map[rune][2]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}
)

func Day15(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	originalMap, err := readTileMap(scanner)
	if err != nil {
		return err
	}

	doubleWideMap := transformToDoubleWideTiles(originalMap)

	var instructions []rune

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		instructions = append(instructions, []rune(line)...)
	}

	for _, tm := range []tileMap{originalMap, doubleWideMap} {
		robotPositions, exists := tm.TypeLocations[robotTile]
		if !exists {
			return errors.New("No robot found in tile map")
		}
		startRobotPos := slices.Collect(maps.Keys(robotPositions))[0]

		curRobotPos := startRobotPos
		for i := 0; i < len(instructions); i++ {
			movementVector := instructionMovements[instructions[i]]

			newRobotPos := [2]int{curRobotPos[0] + movementVector[0], curRobotPos[1] + movementVector[1]}
			newTile, ok := tm.Lookup[newRobotPos]
			if !ok {
				newTile = emptyTile
			}
			switch newTile {
			case wallTile:
				continue
			case emptyTile:
			case boxTile, doubleWideBoxLeftTile, doubleWideBoxRightTile:
				if !moveBoxTiles(&tm, newRobotPos, movementVector) {
					continue
				}
			}
			tm.MoveTiles([][2]int{curRobotPos}, movementVector)
			curRobotPos = newRobotPos
		}

		_, err = tm.WriteTo(os.Stdout)
		if err != nil {
			return err
		}
	}

	var total int
	for boxPos := range originalMap.TypeLocations[boxTile] {
		total += boxPos[0]*100 + boxPos[1]
	}

	fmt.Println("Part 1:", total)

	var total2 int
	for boxPos := range doubleWideMap.TypeLocations[doubleWideBoxLeftTile] {
		total2 += boxPos[0]*100 + boxPos[1]
	}
	fmt.Println("Part 2:", total2)

	return err
}

type tileMap struct {
	Width     int
	Height    int
	EmptyTile rune
	// tile -> [row,col]
	TypeLocations map[rune]map[[2]int]struct{}
	Lookup        map[[2]int]rune
}

func (tm *tileMap) MoveTiles(tilesToMove [][2]int, movementVector [2]int) {
	for i := len(tilesToMove) - 1; i >= 0; i-- {
		oldTilePos := tilesToMove[i]

		tile := tm.Lookup[oldTilePos]
		tm.Lookup[oldTilePos] = tm.EmptyTile
		delete(tm.TypeLocations[tile], oldTilePos)
		tm.TypeLocations[tm.EmptyTile][oldTilePos] = struct{}{}

		newTilePos := [2]int{oldTilePos[0] + movementVector[0], oldTilePos[1] + movementVector[1]}
		tm.Lookup[newTilePos] = tile
		tm.TypeLocations[tile][newTilePos] = struct{}{}
	}
}

func (tm *tileMap) WriteTo(w io.Writer) (n int64, err error) {
	var buf strings.Builder
	for i := 0; i < tm.Height; i++ {
		for j := 0; j < tm.Width; j++ {
			tile, ok := tm.Lookup[[2]int{i, j}]
			if !ok {
				buf.WriteRune(tm.EmptyTile)
				continue
			}
			buf.WriteRune(tile)
		}
		buf.WriteRune('\n')
	}

	return io.Copy(w, strings.NewReader(buf.String()))
}

func (tm *tileMap) Clear() {
	for i := 0; i < tm.Height; i++ {
		for j := 0; j < tm.Width; j++ {
			pos := [2]int{i, j}
			tile, ok := tm.Lookup[pos]
			if ok && tile != tm.EmptyTile {
				delete(tm.TypeLocations, tile)
			}
			tm.SetTile(pos, tm.EmptyTile)
		}
	}
}

func (tm *tileMap) SetTile(pos [2]int, tile rune) {
	if tm.Lookup == nil {
		tm.Lookup = map[[2]int]rune{}
	}
	prevTile, hasPrev := tm.Lookup[pos]
	tm.Lookup[pos] = tile
	if tm.TypeLocations == nil {
		tm.TypeLocations = map[rune]map[[2]int]struct{}{}
	}
	if _, ok := tm.TypeLocations[tile]; !ok {
		tm.TypeLocations[tile] = map[[2]int]struct{}{}
	}
	if hasPrev {
		delete(tm.TypeLocations[prevTile], pos)
	}
	tm.TypeLocations[tile][pos] = struct{}{}
}

func readTileMap(scanner *bufio.Scanner) (tileMap, error) {
	result := tileMap{
		EmptyTile:     emptyTile,
		TypeLocations: map[rune]map[[2]int]struct{}{},
		Lookup:        map[[2]int]rune{},
	}

	i := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if len(result.Lookup) == 0 {
				continue
			}
			result.Height = i
			return result, scanner.Err()
		}

		if result.Width == 0 {
			result.Width = len(line)
		}

		for j, tile := range line {
			if _, exists := result.TypeLocations[tile]; !exists {
				result.TypeLocations[tile] = map[[2]int]struct{}{}
			}

			result.TypeLocations[tile][[2]int{i, j}] = struct{}{}
			result.Lookup[[2]int{i, j}] = tile
		}
		i++
	}
	result.Height = i

	return result, scanner.Err()
}

func transformToDoubleWideTiles(tm tileMap) tileMap {
	result := tileMap{
		EmptyTile:     emptyTile,
		Width:         tm.Width * 2,
		Height:        tm.Height,
		Lookup:        map[[2]int]rune{},
		TypeLocations: map[rune]map[[2]int]struct{}{},
	}

	for i := 0; i < tm.Height; i++ {
		for j := 0; j < tm.Height; j++ {
			tile, ok := tm.Lookup[[2]int{i, j}]
			if !ok {
				tile = emptyTile
			}

			var newTiles [2]rune
			switch tile {
			case robotTile:
				newTiles = [2]rune{robotTile, emptyTile}
			case boxTile:
				newTiles = [2]rune{doubleWideBoxLeftTile, doubleWideBoxRightTile}
			default:
				newTiles = [2]rune{tile, tile}
			}

			for p := range 2 {
				translatedPos := [2]int{i, j*2 + p}
				result.Lookup[translatedPos] = newTiles[p]
				if _, exists := result.TypeLocations[newTiles[p]]; !exists {
					result.TypeLocations[newTiles[p]] = map[[2]int]struct{}{}
				}
				result.TypeLocations[newTiles[p]][translatedPos] = struct{}{}
			}
		}
	}

	return result
}

func moveBoxTiles(tm *tileMap, startBoxPos [2]int, movementVector [2]int) bool {
	upDown := movementVector[0] != 0

	var boxTilePositions [][2]int

	horizontalSearchDeltas := map[rune]int{doubleWideBoxLeftTile: 1, doubleWideBoxRightTile: -1}
	visited := map[[2]int]struct{}{}
	toSearch := [][2]int{startBoxPos}

	for len(toSearch) > 0 {
		var nextSearch [][2]int
		for _, tilePos := range toSearch {
			if _, haveVisited := visited[tilePos]; haveVisited {
				continue
			}
			visited[tilePos] = struct{}{}

			tile, ok := tm.Lookup[tilePos]
			if !ok {
				tile = emptyTile
			}

			switch tile {
			case wallTile:
				return false
			case emptyTile:
			case boxTile:
				boxTilePositions = append(boxTilePositions, tilePos)
				nextSearch = append(nextSearch, [2]int{tilePos[0] + movementVector[0], tilePos[1] + movementVector[1]})
			case doubleWideBoxLeftTile, doubleWideBoxRightTile:
				boxTilePositions = append(boxTilePositions, tilePos)
				nextSearch = append(nextSearch, [2]int{tilePos[0] + movementVector[0], tilePos[1] + movementVector[1]})
				if !upDown {
					continue
				}
				// Add other box tile to search space.
				nextSearch = append(nextSearch, [2]int{tilePos[0], tilePos[1] + horizontalSearchDeltas[tile]})
			}
		}

		toSearch = nextSearch
	}
	tm.MoveTiles(boxTilePositions, movementVector)
	return true
}
