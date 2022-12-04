package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func mustInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func parseRange(s string) [2]int {
	r := strings.Split(s, "-")

	return [2]int{mustInt(r[0]), mustInt(r[1])}
}

func Day4() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day4.txt"))
	if err != nil {
		return err
	}

	var fullyOverlapping int
	var partialOverlapped int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()

		ranges := strings.Split(t, ",")
		r1 := parseRange(ranges[0])
		r2 := parseRange(ranges[1])

		if (r1[0] <= r2[0] && r1[1] >= r2[1]) || (r2[0] <= r1[0] && r2[1] >= r1[1]) {
			fullyOverlapping++
			partialOverlapped++
		} else if (r1[0] >= r2[0] && r1[0] <= r2[1]) || (r1[1] >= r2[0] && r1[1] <= r2[1]) {
			partialOverlapped++
		}
	}

	fmt.Println(fullyOverlapping)
	fmt.Println(partialOverlapped)

	return nil
}
