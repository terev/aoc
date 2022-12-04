package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findWinningPlay(against string) string {
	switch against {
	case "A":
		return "B"
	case "B":
		return "C"
	case "C":
		return "A"
	}
	return ""
}

func findLosingPlay(against string) string {
	switch against {
	case "A":
		return "C"
	case "B":
		return "A"
	case "C":
		return "B"
	}
	return ""
}

func symbolScore(symbol string) int {
	return int(symbol[0] - 64)
}

func playRound(p1, p2 string) int {
	symbolScore := symbolScore(p1)
	if p1 == p2 {
		return 3 + symbolScore
	}

	winningPlay := findWinningPlay(p2)

	if p1 == winningPlay {
		return 6 + symbolScore
	}

	return symbolScore
}

func Day2() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day2.txt"))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	var p1Score int
	var p2Score int
	for scanner.Scan() {
		t := scanner.Text()

		round := strings.Split(t, " ")

		p1Score += playRound(string([]byte(round[1])[0]-23), round[0])

		switch round[1] {
		case "X":
			p2Score += symbolScore(findLosingPlay(round[0]))
		case "Y":
			p2Score += 3 + symbolScore(round[0])
		case "Z":
			p2Score += 6 + symbolScore(findWinningPlay(round[0]))
		}
	}

	fmt.Println(p1Score)
	fmt.Println(p2Score)
	return nil
}
