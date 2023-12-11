package main

import (
	"aoc/util"
	"bufio"
	"io"
	"strings"
)

func Day6(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	raceDurationData := scanner.Text()
	scanner.Scan()
	recordData := scanner.Text()

	raceDurationFields := strings.Fields(raceDurationData)[1:]
	var raceDurations []int
	for _, rd := range raceDurationFields {
		raceDurations = append(raceDurations, util.MustInt(rd))
	}

	p2Duration := util.MustInt(strings.Join(raceDurationFields, ""))

	recordFields := strings.Fields(recordData)[1:]
	var records []int
	for _, r := range recordFields {
		records = append(records, util.MustInt(r))
	}
	p2Record := util.MustInt(strings.Join(recordFields, ""))

	var p1 int
	for race := 0; race < len(raceDurations); race++ {
		var waysToWin int
		for t := 0; t < raceDurations[race]; t++ {
			if t*(raceDurations[race]-t) > records[race] {
				waysToWin++
			}
		}

		if waysToWin > 0 {
			if p1 == 0 {
				p1 = waysToWin
			} else {
				p1 *= waysToWin
			}
		}
	}

	var p2 int
	for t := 0; t < p2Duration; t++ {
		if t*(p2Duration-t) > p2Record {
			p2++
		}
	}

	return p1, p2, nil
}
