package main

import (
	"aoc/input"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Solution struct {}

func (s *Solution) Date() (int, int) {
	return 2018, 3
}

type claim struct {
	claimID int
	x int
	y int
	width int
	height int
}

func (s *Solution) Execute(input []byte) error {
	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	//var area map[int]map[int]int

	var fabric[1000][1000]struct {
		overlap int
		claims map[int]int
	}

	var claimers = make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		newclaim := &claim{}

		_, err := fmt.Fscanf(strings.NewReader(line), "#%d @ %d,%d: %dx%d", &newclaim.claimID, &newclaim.x, &newclaim.y, &newclaim.width, &newclaim.height)

		if err != nil {
			return err
		}

		for j := newclaim.y; j < newclaim.y + newclaim.height; j++ {
			for i := newclaim.x; i < newclaim.x + newclaim.width; i++ {
				fabric[j][i].overlap++
				if len(fabric[j][i].claims) > 0 {
					fabric[j][i].claims[newclaim.claimID] = 1
				} else {
					fabric[j][i].claims = make(map[int]int)
					fabric[j][i].claims[newclaim.claimID] = 1
				}
			}
		}

		claimers[newclaim.claimID] = 0
	}

	var total int
	for j := 0; j < len(fabric); j++ {
		for i := 0; i < len(fabric[j]); i++ {
			if fabric[j][i].overlap > 1 {
				total++

				for claimer := range fabric[j][i].claims {
					claimers[claimer]++
				}
			}
		}
	}

	fmt.Println(total)

	for claimid, overlap := range claimers {
		if overlap == 0 {
			fmt.Println(claimid)
		}
	}

	return nil
}


func main() {
	s := Solution{}
	in, err := input.GetInput(s.Date())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = s.Execute(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}