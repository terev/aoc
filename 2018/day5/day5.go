package main

import (
	"aoc/input"
	"bytes"
	"fmt"
	"math"
	"os"
)

type Solution struct{}

func (s *Solution) Date() (int, int) {
	return 2018, 5
}

func reactPolymer(in []byte) []byte {
	inputCopy := make([]byte, len(in))
	copy(inputCopy, in)

	var prev byte
	var i int
	var c byte
	for i = 0; i < len(inputCopy); i++ {
		c = inputCopy[i]

		if prev > 0 && c == prev^0x20 {
			inputCopy = append(inputCopy[:i-1], inputCopy[i+1:]...)
			i = int(math.Max(float64(i-2), 0))
			c = inputCopy[i]
		}

		prev = c
	}

	return inputCopy
}

func (s *Solution) Execute(input []byte) error {
	// rawr
	input = bytes.TrimSpace(input)

	result := reactPolymer(input)
	fmt.Println("Part 1:")
	fmt.Println(string(result))
	fmt.Println(len(result))

	var least int = -1
	for i := 0; i < 25; i ++ {
		inputCopy := make([]byte, len(input))
		copy(inputCopy, input)

		inputCopy = bytes.Replace(inputCopy, []byte{byte(i + 65)}, []byte{}, -1)
		inputCopy = bytes.Replace(inputCopy, []byte{byte(i + 97)}, []byte{}, -1)


		fullyReacted := reactPolymer(inputCopy)

		if least == -1 || len(fullyReacted) < least {
			least = len(fullyReacted)
		}
	}

	fmt.Println("Part 2:")
	fmt.Println(least)

	return nil
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
