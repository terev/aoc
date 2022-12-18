package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func showScan(scan [][]byte) {
	for i := 0; i < len(scan); i++ {
		for j := 0; j < len(scan[i]); j++ {
			if scan[i][j] != 0 {
				fmt.Printf("%c", scan[i][j])
			} else {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Println()
	}
}

func Day14() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day14.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	var xbounds, ybounds [2]int
	xbounds[0] = 5000
	ybounds[0] = 5000
	var walls [][][2]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		structure := scanner.Text()
		transitions := strings.Split(structure, "->")

		var trace [][2]int

		for _, transition := range transitions {
			coordsRaw := strings.Split(strings.TrimSpace(transition), ",")
			coords := [2]int{util.MustInt(coordsRaw[0]), util.MustInt(coordsRaw[1])}

			if coords[0] < xbounds[0] {
				xbounds[0] = coords[0]
			}

			if coords[0] > xbounds[1] {
				xbounds[1] = coords[0]
			}

			if coords[1] < ybounds[0] {
				ybounds[0] = coords[1]
			}

			if coords[1] > ybounds[1] {
				ybounds[1] = coords[1]
			}
			trace = append(trace, coords)
		}

		walls = append(walls, trace)
	}

	fmt.Println(xbounds, ybounds)

	var (
		width  = xbounds[1] - xbounds[0] + 3
		height = ybounds[1] + 2
	)

	xbounds[0]--

	scan := make([][]byte, height)
	for i := range scan {
		scan[i] = make([]byte, width)
	}

	fmt.Println(width, height)

	for _, trace := range walls {
		var prev = trace[0]
		for i := 1; i < len(trace); i++ {
			var increment [2]int

			if trace[i][0] != prev[0] {
				if trace[i][0] > prev[0] {
					increment = [2]int{1, 0}
				} else {
					increment = [2]int{-1, 0}
				}
			} else {
				if trace[i][1] > prev[1] {
					increment = [2]int{0, 1}
				} else {
					increment = [2]int{0, -1}
				}
			}

			for prev[0] != trace[i][0] || prev[1] != trace[i][1] {
				scan[prev[1]][prev[0]-xbounds[0]] = '#'
				prev[0] += increment[0]
				prev[1] += increment[1]
				scan[prev[1]][prev[0]-xbounds[0]] = '#'
			}

			prev = trace[i]
		}
	}

	showScan(scan)

	origin := [2]int{500 - xbounds[0], 0}

	sandPos := [2]int{origin[0], origin[1]}
	sand := 1
	for sandPos[1] < ybounds[1] {
		if scan[sandPos[1]+1][sandPos[0]] == 0 {
			sandPos[1]++
		} else if scan[sandPos[1]+1][sandPos[0]-1] == 0 {
			sandPos[0]--
			sandPos[1]++
		} else if scan[sandPos[1]+1][sandPos[0]+1] == 0 {
			sandPos[0]++
			sandPos[1]++
		} else {
			scan[sandPos[1]][sandPos[0]] = 'o'
			sandPos = [2]int{origin[0], origin[1]}
			sand++
		}
	}

	showScan(scan)

	fmt.Println(sand - 1)

	for {
		if sandPos[0]+1 >= width {
			for i := range scan {
				scan[i] = append(scan[i], 0)
			}
			width++
		} else if sandPos[0]-1 < 0 {
			for i := range scan {
				scan[i] = append([]byte{0}, scan[i]...)
			}
			origin[0]++
			width++
			sandPos[0]++
		}

		if sandPos[1] >= ybounds[1]+1 {
			scan[sandPos[1]][sandPos[0]] = 'o'
			sandPos = [2]int{origin[0], origin[1]}
			sand++
		} else if scan[sandPos[1]+1][sandPos[0]] == 0 {
			sandPos[1]++
		} else if scan[sandPos[1]+1][sandPos[0]-1] == 0 {
			sandPos[0]--
			sandPos[1]++
		} else if scan[sandPos[1]+1][sandPos[0]+1] == 0 {
			sandPos[0]++
			sandPos[1]++
		} else if sandPos[0] == origin[0] && sandPos[1] == origin[1] {
			break
		} else {
			scan[sandPos[1]][sandPos[0]] = 'o'
			sandPos = [2]int{origin[0], origin[1]}
			sand++
		}
	}

	showScan(scan)
	fmt.Println(sand)

	return nil
}
