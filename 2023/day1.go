package main

import (
	"bufio"
	"io"
	"strings"
)

func Day1(r io.Reader) (int, int, error) {

	search := []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
	search2 := append(search,
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine")

	scanner := bufio.NewScanner(r)

	var sum, sum2 int
	for scanner.Scan() {
		line := scanner.Text()
		var (
			first         = len(line)
			last          = 0
			firstn, lastn int
		)
		for i, s := range search {
			l := strings.Index(line, s)
			r := strings.LastIndex(line, s)
			if l != -1 && l <= first {
				first = l
				firstn = i
			}
			if r != -1 && r >= last {
				last = r
				lastn = i
			}
		}

		sum += firstn*10 + lastn

		first = len(line)
		last = 0

		for i, s := range search2 {
			l := strings.Index(line, s)
			r := strings.LastIndex(line, s)

			if l != -1 && l <= first {
				first = l
				firstn = i
			}
			if r != -1 && r >= last {
				last = r
				lastn = i
			}
		}

		sum2 += (firstn%10)*10 + lastn%10
	}

	return sum, sum2, nil
}
