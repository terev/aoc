package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Day10() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day10.txt"))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	cycle := 1
	register := 1
	instruction := []string{}
	instructionTime := 0

	var sum int

	for {
		if instructionTime == 0 {
			if len(instruction) > 0 && instruction[0] == "addx" {
				register += util.MustInt(instruction[1])
			}

			if !scanner.Scan() {
				break
			}

			instruction = strings.Split(scanner.Text(), " ")
			switch instruction[0] {
			case "noop":
				instructionTime = 1
			case "addx":
				instructionTime = 2
			}
		}

		if (cycle-20)%40 == 0 {
			sum += cycle * register
		}

		pixelPos := (cycle - 1) % 40

		if pixelPos == register || pixelPos == register-1 || pixelPos == register+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if cycle%40 == 0 {
			fmt.Println()
		}

		instructionTime--
		cycle++
	}

	fmt.Println()
	fmt.Println(sum)

	return nil
}
