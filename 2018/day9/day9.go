package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

type Solution struct{}

func (s *Solution) Date() (int, int) {
	return 2018, 9
}

func (s *Solution) Execute(input Input) error {
	scanner := bufio.NewScanner(bytes.NewReader(input.data))

	var (
		playerCount int
		lastMarble  int
	)

	if scanner.Scan() {
		if _, err := fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &playerCount, &lastMarble); err != nil {
			return err
		}
	} else {
		return scanner.Err()
	}

	//lastMarble *= 100

	var playerScores = make([]int, playerCount)

	var board = make([]int, lastMarble)
	var size = 1
	var currentIdx = 0

	//printBoard(board, currentMarbleIdx)
	for marbleIdx := 1; marbleIdx < lastMarble; marbleIdx++ {
		if marbleIdx % 23 == 0 {
			playerScores[marbleIdx%playerCount] += marbleIdx
		} else {
			board[int(math.Abs(float64(currentIdx - size % lastMarble)))] = board[currentIdx]
			currentIdx = int(math.Abs(float64((currentIdx - 1) % lastMarble)))
			board[int(math.Abs(float64(currentIdx - size % lastMarble)))] = board[currentIdx]
			board[currentIdx] = marbleIdx
			size ++
		}
		printBoard(board, currentIdx)
	}

	var topscore = playerScores[0]
	for i := 1; i < playerCount; i ++ {
		if playerScores[i] > topscore {
			topscore = playerScores[i]
		}
	}

	fmt.Println(topscore)

	return nil
}

func printBoard(board []int, currentMarbleIdx int) {
	for j, e := range board {
		if j == currentMarbleIdx {
			fmt.Printf(" (%d) ", e)
		} else {
			fmt.Printf(" %d ", e)
		}
	}
	fmt.Println()
}

var samples = []string{
	`9 players; last marble is worth 25 points`,
	//`10 players; last marble is worth 1618 points`,
	//`30 players; last marble is worth 5807 points`,
}

type Input struct {
	data []byte
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	solution := Solution{}

	fmt.Println("Sample:")

	for _, sample := range samples {
		checkError(
			solution.Execute(Input{
				data: []byte(sample),
			}),
		)
	}

	//in, err := input.GetInput(solution.Date())
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println("Real:")
	//checkError(
	//	solution.Execute(Input{
	//		data: in,
	//	}),
	//)
}
