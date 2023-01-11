package main

import (
	"aoc/input"
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"os"
)

type Solution struct{}

func (s *Solution) Date() (int, int) {
	return 2018, 8
}

type node struct {
	childrenCount int
	metadataCount int

	children []*node
	metadata []int
}

func scanInt(scanner *bufio.Scanner) (int, bool) {
	ok := scanner.Scan()
	if !ok {
		return 0, ok
	}

	return cast.ToInt(scanner.Text()), true
}

func (s *Solution) Execute(input Input) error {
	scanner := bufio.NewScanner(bytes.NewReader(input.data))

	scanner.Split(bufio.ScanWords)

	var tree = scanTree(scanner)

	fmt.Println("Part 1:")
	fmt.Println(sumTree(tree))
	fmt.Println("Part 2:")
	fmt.Println(sumTree2(tree))

	return scanner.Err()
}

func sumTree(current *node) int {
	if current == nil {
		return 0
	} else {
		var sum int
		for _, entry := range current.metadata {
			sum += entry
		}

		for _, child := range current.children {
			sum += sumTree(child)
		}

		return sum
	}
}

func sumTree2(current *node) int {
	if current == nil {
		return 0
	} else {
		var sum int

		if len(current.children) == 0 {
			for _, entry := range current.metadata {
				sum += entry
			}
		} else {
			for _, entry := range current.metadata {
				if (entry - 1) < len(current.children) {
					sum += sumTree2(current.children[entry-1])
				}
			}
		}

		return sum
	}
}

func scanTree(scanner *bufio.Scanner) *node {
	newNode := &node{}

	var ok bool
	if newNode.childrenCount, ok = scanInt(scanner); !ok {
		return nil
	}

	if newNode.metadataCount, ok = scanInt(scanner); !ok {
		return nil
	}

	for i := 0; i < newNode.childrenCount; i++ {
		newChild := scanTree(scanner)
		if newChild == nil {
			break
		}

		newNode.children = append(newNode.children, newChild)
	}

	for i := 0; i < newNode.metadataCount; i++ {
		newInt, ok := scanInt(scanner)
		if !ok {
			return nil
		}

		newNode.metadata = append(newNode.metadata, newInt)
	}

	return newNode
}

var sample = `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`

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
	checkError(
		solution.Execute(Input{
			data: []byte(sample),
		}),
	)

	in, err := input.GetInput(solution.Date())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Real:")
	checkError(
		solution.Execute(Input{
			data: in,
		}),
	)
}
