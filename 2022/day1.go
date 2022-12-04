package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func Day1() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day1.txt"))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	var sum int

	var calories []int

	for scanner.Scan() {
		t := scanner.Text()

		if t == "" {
			calories = append(calories, sum)
			sum = 0

			continue
		}

		c, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return err
		}

		sum += int(c)
	}

	calories = append(calories, sum)

	sort.Ints(calories)

	fmt.Println(calories[len(calories)-1])

	var topThree int
	for _, c := range calories[len(calories)-3:] {
		fmt.Println(c)
		topThree += c
	}

	fmt.Println(topThree)

	return nil
}
