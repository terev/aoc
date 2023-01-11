package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

type Solution struct {}

func (s *Solution) Date() (int, int) {
	return 2018, 1
}

func (s *Solution) Execute(input []byte) error {
	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	frequencies := make(map[int]int)

	var (
		total int
		loops int
		found bool
	)

	for !found {
		for scanner.Scan() {
			frequencies[total] += 1
			if frequencies[total] >= 2 {
				fmt.Printf("First reached %d twice\n", total)
				found = true
				if loops > 0 {
					break
				}
			}

			change, err := cast.ToIntE(scanner.Text())
			if err != nil {
				return err
			}

			total += change
		}
		if !found {
			scanner = bufio.NewScanner(strings.NewReader(string(input)))
		}

		if loops == 0 {
			fmt.Printf("Resulting frequency: %d\n", total)
		}
		loops ++
	}

	return nil
}



