package main

import (
	"aoc/util"
	"bytes"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var shapesRaw = []string{
	`####`,
	`.#.
###
.#.`,
	`..#
..#
###`,
	`#
#
#
#`,
	`##
##`,
}

var shapeBounds = [][2]int{
	{4, 1},
	{3, 3},
	{3, 3},
	{1, 4},
	{2, 2},
}

var shapes [][]string

func init() {
	for i := 0; i < len(shapesRaw); i++ {
		shapes = append(shapes, strings.Split(shapesRaw[i], "\n"))
	}
}

func checkCollision(chamber [][7]byte, shape []string, pos [2]int) bool {
	for i := 0; i < len(shape); i++ {
		if pos[0]+i < 0 || pos[0]+i > len(chamber) {
			continue
		}
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == '#' {
				if pos[0]+i == len(chamber) || (chamber[pos[0]+i][pos[1]+j] == '#') {
					return true
				}
			}
		}
	}

	return false
}

func placeRock(chamber [][7]byte, shape []string, pos [2]int) [][7]byte {
	if pos[0] < 0 {
		diff := util.Abs(pos[0])
		newRows := make([][7]byte, diff)
		chamber = slices.Insert(chamber, 0, newRows...)
		pos[0] += diff
	}

	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == '#' {
				chamber[pos[0]+i][pos[1]+j] = '#'
			}
		}
	}

	return chamber
}

func showChamber(chamber [][7]byte) {
	for i := 0; i < len(chamber); i++ {
		for j := 0; j < len(chamber[i]); j++ {
			if chamber[i][j] == '#' {
				fmt.Printf("%c", '#')
			} else {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Println()
	}
}

func towerHeight(rocksToDrop int, jets []byte) int {
	var (
		curJet       int
		curShape     int
		rocksDropped int
		chamber      = make([][7]byte, 1)
		rockPos      = [2]int{-3, 2}
	)

	for rocksDropped < rocksToDrop {
		switch jets[curJet] {
		case '<':
			if rockPos[1] > 0 && !checkCollision(chamber, shapes[curShape], [2]int{rockPos[0], rockPos[1] - 1}) {
				rockPos[1]--
			}
		case '>':
			if rockPos[1]+shapeBounds[curShape][0] < 7 && !checkCollision(chamber, shapes[curShape], [2]int{rockPos[0], rockPos[1] + 1}) {
				rockPos[1]++
			}
		}

		curJet = (curJet + 1) % len(jets)

		if !checkCollision(chamber, shapes[curShape], [2]int{rockPos[0] + 1, rockPos[1]}) {
			rockPos[0]++
			continue
		}

		chamber = placeRock(chamber, shapes[curShape], rockPos)

		rocksDropped++
		curShape = rocksDropped % len(shapes)
		rockPos = [2]int{-3 - shapeBounds[curShape][1], 2}

		if rocksDropped%1000000 == 0 {
			fmt.Println(rocksDropped)
		}
	}

	return len(chamber)
}

func Day17() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day17.txt"))
	if err != nil {
		return err
	}

	jets, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	f.Close()

	jets = bytes.TrimSpace(jets)

	fmt.Println(towerHeight(2022, jets))
	fmt.Println(towerHeight(1000000000000, jets))

	return nil
}
