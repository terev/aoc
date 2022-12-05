package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Day5() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day5.txt"))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	var stacks [][]byte
	var legendFound bool

	for scanner.Scan() {
		t := scanner.Text()

		var j int
		for i := 0; i < len(t); i += 4 {
			element := t[i : i+3]
			if element == " 1 " {
				legendFound = true
				break
			}

			item := strings.Trim(element, "[]")[0]

			if j+1 > len(stacks) {
				stacks = append(stacks, []byte{})
			}

			if item != ' ' {
				stacks[j] = append(stacks[j], item)
			}

			j++
		}

		if legendFound {
			break
		}
	}

	// Make a copy of parsed stacks for part2
	part2stacks := make([][]byte, len(stacks))
	for i := range stacks {
		part2stacks[i] = make([]byte, len(stacks[i]))
		copy(part2stacks[i], stacks[i])
	}

	for scanner.Scan() {
		t := scanner.Text()

		// Skip empty lines
		if t == "" {
			continue
		}

		instruction := strings.Split(t, " ")
		toMove := util.MustInt(instruction[1])
		from := util.MustInt(instruction[3]) - 1
		to := util.MustInt(instruction[5]) - 1

		for i := 0; i < toMove; i++ {
			item := stacks[from][0]
			if len(stacks[from]) == 1 {
				stacks[from] = []byte{}
			} else {
				stacks[from] = stacks[from][1:]
			}

			stacks[to] = append([]byte{item}, stacks[to]...)
		}

		items := make([]byte, toMove+len(part2stacks[to]))
		copy(items, part2stacks[from][:toMove])
		if len(part2stacks[from]) == toMove {
			part2stacks[from] = []byte{}
		} else {
			part2stacks[from] = part2stacks[from][toMove:]
		}

		copy(items[toMove:], part2stacks[to])
		part2stacks[to] = items
	}

	var message string
	var part2message string
	for i := 0; i < len(stacks); i++ {
		if len(stacks[i]) > 0 {
			message += string(stacks[i][0])
		}

		if len(part2stacks[i]) > 0 {
			part2message += string(part2stacks[i][0])
		}
	}

	fmt.Println(message)
	fmt.Println(part2message)

	return nil
}
