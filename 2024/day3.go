package main

import (
	"fmt"
	"io"
	"regexp"

	"aoc/util"
)

var mulRe = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var mulToggleableRe = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

func Day3P1(r io.Reader) error {
	in, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	inText := string(in)

	matches := mulRe.FindAllStringSubmatch(inText, -1)

	var tot int
	for _, match := range matches {
		tot += util.MustInt(match[1]) * util.MustInt(match[2])
	}

	fmt.Println("P1:", tot)
	return nil
}

func Day3P2(r io.Reader) error {
	in, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	inText := string(in)

	matches := mulToggleableRe.FindAllStringSubmatch(inText, -1)

	mulEnabled := true
	tot2 := 0
	for _, match := range matches {
		switch match[0] {
		case "don't()":
			mulEnabled = false
			continue
		case "do()":
			mulEnabled = true
		default:
			if mulEnabled {
				tot2 += util.MustInt(match[1]) * util.MustInt(match[2])
			}
		}
	}
	fmt.Println("P2:", tot2)
	return nil
}
