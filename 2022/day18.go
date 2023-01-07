package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"path/filepath"
	"strings"
)

type cube struct {
	corners map[[3]int]struct{}
}

func Day18() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day18.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	corners := map[[3]int][]*cube{}

	potentialAirPockets := map[[3]int]int{}

	origins := map[[3]int]struct{}{}

	var openFaces int

	for scanner.Scan() {
		originRaw := scanner.Text()

		originParts := strings.Split(originRaw, ",")

		origin := [3]int{util.MustInt(originParts[0]), util.MustInt(originParts[1]), util.MustInt(originParts[2])}
		origins[origin] = struct{}{}

		for i := 0; i < 3; i++ {
			potentialAirPocket := origin
			potentialAirPocket[i] += 1
			potentialAirPockets[potentialAirPocket]++
			potentialAirPocket = origin
			potentialAirPocket[i] += -1
			potentialAirPockets[potentialAirPocket]++
		}

		openFaces += 6

		var newCube = &cube{}

		var faces [3][2][][3]int

		for i := 0; i < 8; i++ {
			var translated [3]int
			var cornerFaces [][2]int
			faceMask := 1
			for j := 0; j < 3; j++ {
				translation := (i & faceMask) >> j
				cornerFaces = append(cornerFaces, [2]int{j, translation})
				translated[j] = origin[j] + translation
				faceMask <<= 1
			}

			corners[translated] = append(corners[translated], newCube)

			for _, f := range cornerFaces {
				faces[f[0]][f[1]] = append(faces[f[0]][f[1]], translated)
			}
		}

		for _, dim := range faces {
			for _, f := range dim {
				var potentialCubes []*cube
				for i, cornerCoords := range f {
					c, exists := corners[cornerCoords]
					if !exists {
						break
					}

					if i == 0 {
						for _, sharedCorner := range c {
							if sharedCorner != newCube {
								potentialCubes = append(potentialCubes, sharedCorner)
							}
						}
						continue
					}

					for j := 0; j < len(potentialCubes); j++ {
						if !slices.Contains(c, potentialCubes[j]) {
							potentialCubes = slices.Delete(potentialCubes, j, j+1)
							j--
						}
					}

					if len(potentialCubes) == 0 {
						break
					}
				}

				if len(potentialCubes) > 1 {
					panic("That shouldnt happen")
				} else if len(potentialCubes) == 1 {
					openFaces -= 2
				}
			}
		}
	}

	fmt.Println(openFaces)

	for p, matching := range potentialAirPockets {
		if matching == 6 {
			if _, isOrigin := origins[p]; !isOrigin {
				fmt.Println(p)
				openFaces -= 6
			}
		}
	}

	fmt.Println(openFaces)

	return nil
}
