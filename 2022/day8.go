package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Day8() error {
	var forest [][]int

	f, err := os.Open(filepath.Join(util.Cwd(), "day8.txt"))
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		row := make([]int, len(t))

		for i, height := range t {
			row[i] = util.MustInt(string(height))
		}

		forest = append(forest, row)
	}

	visibilities := make([][]int, len(forest))
	for i := range forest {
		visibilities[i] = make([]int, len(forest[i]))
	}

	var totalVisible int

	for i := range forest {
		var (
			maxLeft  = -1
			maxRight = -1
		)

		rowSize := len(forest[i])
		for j := range forest[i] {
			if forest[i][j] > maxLeft {
				maxLeft = forest[i][j]
				visibilities[i][j]++
				if visibilities[i][j] == 1 {
					totalVisible++
				}
			}

			fromRight := rowSize - j - 1
			if forest[i][fromRight] > maxRight {
				maxRight = forest[i][fromRight]
				visibilities[i][fromRight]++
				if visibilities[i][fromRight] == 1 {
					totalVisible++
				}
			}
		}
	}

	forestHeight := len(forest)

	for j := 0; j < len(forest[0]); j++ {
		var (
			maxTop    = -1
			maxBottom = -1
		)
		for i := range forest {
			if forest[i][j] > maxTop {
				maxTop = forest[i][j]
				visibilities[i][j]++
				if visibilities[i][j] == 1 {
					totalVisible++
				}
			}

			fromBottom := forestHeight - i - 1
			if forest[fromBottom][j] > maxBottom {
				maxBottom = forest[fromBottom][j]
				visibilities[fromBottom][j]++
				if visibilities[fromBottom][j] == 1 {
					totalVisible++
				}
			}
		}
	}

	fmt.Println(totalVisible)

	forestWidth := len(forest[0])
	bestScore := 0
	for i := range forest {
		for j := range forest[i] {
			directionScores := [4]struct {
				blocked bool
				score   int
			}{}
			allBlocked := false
			k := 1
			for !allBlocked {
				if !directionScores[0].blocked {
					if i-k < 0 {
						directionScores[0].blocked = true
					} else {
						directionScores[0].score++
						if forest[i-k][j] >= forest[i][j] {
							directionScores[0].blocked = true
						}
					}
				}

				if !directionScores[2].blocked {
					if i+k > forestHeight-1 {
						directionScores[2].blocked = true
					} else {
						directionScores[2].score++
						if forest[i+k][j] >= forest[i][j] {
							directionScores[2].blocked = true
						}
					}
				}

				if !directionScores[1].blocked {
					if j+k > forestWidth-1 {
						directionScores[1].blocked = true
					} else {
						directionScores[1].score++
						if forest[i][j+k] >= forest[i][j] {
							directionScores[1].blocked = true
						}
					}
				}

				if !directionScores[3].blocked {
					if j-k < 0 {
						directionScores[3].blocked = true
					} else {
						directionScores[3].score++
						if forest[i][j-k] >= forest[i][j] {
							directionScores[3].blocked = true
						}
					}
				}

				allBlocked = directionScores[0].blocked && directionScores[1].blocked &&
					directionScores[2].blocked && directionScores[3].blocked
				k++
			}
			var totalScore int = 1
			for _, dir := range directionScores {
				totalScore *= dir.score
			}

			if totalScore > bestScore {
				bestScore = totalScore
			}
		}
	}

	fmt.Println(bestScore)

	return nil
}
