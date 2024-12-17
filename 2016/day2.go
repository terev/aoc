package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

func Day2(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	var buttons []int
	pos := [2]int{1, 1}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		for _, c := range line {
			switch unicode.ToUpper(c) {
			case 'U':
				pos[1] = max(pos[1]-1, 0)
			case 'D':
				pos[1] = min(pos[1]+1, 2)
			case 'L':
				pos[0] = max(pos[0]-1, 0)
			case 'R':
				pos[0] = min(pos[0]+1, 2)
			}
		}
		buttons = append(buttons, (pos[0]+(pos[1]*3))+1)
	}

	fmt.Println(buttons)
	return nil
}
