package main

import (
	"aoc/util"
	"bufio"
	"io"
	"slices"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

func Day15(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	var p1 int
	var boxes [256][]lens
	for scanner.Scan() {
		line := scanner.Text()

		steps := strings.Split(line, ",")

		for _, step := range steps {
			h := hash(step)
			p1 += h

			label, operation, parameter := parseStep(step)
			box := hash(label)
			idx := slices.IndexFunc(boxes[box], func(l lens) bool {
				return l.label == label
			})
			switch operation {
			case 0:
				if idx == -1 {
					boxes[box] = append(boxes[box], lens{
						label:       label,
						focalLength: parameter,
					})
				} else {
					boxes[box][idx].focalLength = parameter
				}
			case 1:
				if idx != -1 {
					boxes[box] = slices.Delete(boxes[box], idx, idx+1)
				}
			}
		}
	}

	var p2 int
	for i, box := range boxes {
		for j, storedLens := range box {
			p2 += (1 + i) * (1 + j) * storedLens.focalLength
		}
	}

	return p1, p2, nil
}

func hash(s string) int {
	var result int
	for _, c := range s {
		result += int(c)
		result *= 17
		result %= 256
	}

	return result
}

func parseStep(step string) (string, uint8, int) {
	for i := 0; i < len(step); i++ {
		switch step[i] {
		case '=':
			return step[:i], 0, util.MustInt(step[i+1:])
		case '-':
			return step[:i], 1, 0
		}
	}

	panic("invalid step")
}
