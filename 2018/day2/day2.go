package main

import (
	"aoc/input"
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"os"
	"strings"
)

type Solution struct {}

func (s *Solution) Date() (int, int) {
	return 2018, 2
}

func (s *Solution) Execute(input []byte) error {
	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	total2 := 0
	total3 := 0

	for scanner.Scan() {
		line := scanner.Text()

		counts := make(map[rune] int)
		for _, c := range line {
			counts[c]++
		}

		var (
			has2 bool
			has3 bool
		)

		for _, count := range counts {
			if count == 2 {
				has2 = true
			} else if count == 3 {
				has3 = true
			}

			if has2 && has3 {
				break
			}
		}

		total2 += cast.ToInt(has2)
		total3 += cast.ToInt(has3)
	}

	fmt.Println("Part 1: ", total2 * total3)

	scanner = bufio.NewScanner(strings.NewReader(string(input)))

	var prev []string
	for scanner.Scan() {
		line := scanner.Text()

		for _, p := range prev {
			diff := differWithinConstraint(1, line, p)

			if diff >= 0 && diff <= 1 {
				fmt.Println("Part 2: ", common(line, p))
				return nil
			}
		}

		prev = append(prev, line)
	}

	return nil
}

func differWithinConstraint(max int, s1 string, s2 string) int {
	var diff int

	for i := 0; i < len(s1); i ++ {
		if s1[i] != s2[i] {
			diff++
		}

		if diff > max {
			return -1
		}
	}

	return diff
}

func common(s1 string, s2 string) string {
	var common string

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			common += string(s1[i])
		}
	}

	return common
}

func main() {
	s := Solution{}
	in, err := input.GetInput(s.Date())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = s.Execute(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}