package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Day20() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day20.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	var enc []int

	var i int

	indexLookup := map[int]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		n := util.MustInt(t)

		indexLookup[n] = i

		enc = append(enc, n)
		i++
	}

	encLength := len(enc)

	dec := make([]int, encLength)

	copy(dec, enc)

	for _, moves := range enc {
		initialNewPosition := (indexLookup[moves] + moves) % encLength

		if initialNewPosition == indexLookup[moves] {
			continue
		}

		replacement := moves
		newPosition := initialNewPosition

		fmt.Println(replacement, newPosition)

		var toReplace int
		for {
			toReplace = dec[newPosition]
			dec[newPosition] = replacement
			indexLookup[replacement] = newPosition

			newPosition = (newPosition + 1) % encLength
			if newPosition == initialNewPosition {
				break
			}

			replacement = toReplace
		}
		fmt.Println(dec)
	}

	fmt.Println(dec)

	return nil
}
