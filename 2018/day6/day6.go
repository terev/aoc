package main

import (
	"aoc/input"
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

type Solution struct{}

func (s *Solution) Date() (int, int) {
	return 2018, 6
}

func manhattan(p1 [2]int, p2 [2]int) int {
	return int(math.Abs(float64(p2[0]-p1[0])) + math.Abs(float64(p2[1]-p1[1])))
}

func (s *Solution) Execute(input Input) error {
	// rawr
	scanner := bufio.NewScanner(bytes.NewReader(input.raw))

	var (
		minx int
		miny int
		maxx int
		maxy int
	)

	var distanceCap = input.fields["distance_cap"].(int)

	var coords [][2]int

	for scanner.Scan() {
		coordsS := scanner.Text()

		var (
			x int
			y int
		)

		_, err := fmt.Sscanf(coordsS, "%d,%d", &x, &y)

		if err != nil {
			return err
		}

		coords = append(coords, [2]int{x, y})

		if minx == 0 || x < minx {
			minx = x
		}

		if miny == 0 || y < miny {
			miny = y
		}

		if maxx == 0 || x > maxx {
			maxx = x
		}

		if maxy == 0 || y > maxy {
			maxy = y
		}
	}

	var regionSize int
	var grid = make([][]map[int][]int, (maxy-miny)+1)

	for i := 0; i < len(grid); i ++ {
		grid[i] = make([]map[int][]int, (maxx-minx)+1)
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = make(map[int][]int)
		}
	}

	var candidateRegionSizes = make(map[int]int)
	for i := minx; i < maxx; i++ {
		for j := miny; j < maxy; j++ {

			var total int
			var best = -1
			var bestdistance = -1
			for id, coord := range coords {
				distanceTo := manhattan([2]int{i, j}, coord)

				total += distanceTo

				if bestdistance == -1 || distanceTo < bestdistance {
					bestdistance = distanceTo
					best = id
				} else if distanceTo == bestdistance && best != -1 {
					best = -1
				}
			}

			if best != -1 {
				candidateRegionSizes[best]++
			}

			if total < distanceCap {
				regionSize++
			}
		}
	}

	var maxSize = -1

	for id, coord := range coords {
		if coord[0] > minx && coord[0] < maxx && coord[1] > miny && coord[1] < maxy {
			if maxSize == -1 || candidateRegionSizes[id] > maxSize {
				maxSize = candidateRegionSizes[id]
			}
		}
	}

	fmt.Println("Max non-infinite:", maxSize)
	fmt.Println("Max < 10000:", regionSize)

	return nil
}

var sample = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

type Input struct {
	raw    []byte
	fields map[string]interface{}
}

func main() {
	s := Solution{}
	in, err := input.GetInput(s.Date())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Sample:")
	err = s.Execute(Input{
		raw: []byte(sample),
		fields: map[string]interface{}{
			"distance_cap": 32,
		},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Real:")
	err = s.Execute(Input{
		raw: in,
		fields: map[string]interface{}{
			"distance_cap": 10000,
		},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
