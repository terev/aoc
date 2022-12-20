package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func sign(a int) int {
	if a == 0 {
		return 0
	}
	if a > 0 {
		return 1
	}
	return -1
}

func simulateKnots(nKnots int) (int, error) {
	knots := make([][2]int, nKnots)

	f, err := os.Open(filepath.Join(util.Cwd(), "day9.txt"))
	if err != nil {
		return 0, err
	}

	uniquePositions := make(map[[2]int]struct{})
	uniquePositions[[2]int{0, 0}] = struct{}{}

	var totalPositions int = 1

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")

		direction := instruction[0]
		increment := util.MustInt(instruction[1])
		var movementVector [2]int

		switch direction {
		case "L":
			movementVector = [2]int{-1, 0}
		case "R":
			movementVector = [2]int{1, 0}
		case "U":
			movementVector = [2]int{0, 1}
		case "D":
			movementVector = [2]int{0, -1}
		}

		for i := 0; i < increment; i++ {
			for j := 0; j < nKnots-1; j++ {
				head := &knots[j]
				tail := &knots[j+1]

				if j == 0 {
					head[0] += movementVector[0]
					head[1] += movementVector[1]
				}

				if head[1] == tail[1] && util.Abs(head[0]-tail[0]) > 1 {
					tail[0] += sign(head[0] - tail[0])
				} else if head[0] == tail[0] && util.Abs(head[1]-tail[1]) > 1 {
					tail[1] += sign(head[1] - tail[1])
				} else if head[0] != tail[0] && head[1] != tail[1] && (util.Abs(head[0]-tail[0]) > 1 || util.Abs(head[1]-tail[1]) > 1) {
					tail[0] += sign(head[0] - tail[0])
					tail[1] += sign(head[1] - tail[1])
				}

				if j == nKnots-2 {
					if _, exists := uniquePositions[*tail]; !exists {
						totalPositions++
						uniquePositions[*tail] = struct{}{}
					}
				}
			}
		}
	}

	return totalPositions, nil
}

func Day9() error {
	totalPositions, err := simulateKnots(2)
	if err != nil {
		return err
	}

	fmt.Println(totalPositions)

	totalPositions2, err := simulateKnots(10)
	if err != nil {
		return err
	}

	fmt.Println(totalPositions2)

	return nil
}
