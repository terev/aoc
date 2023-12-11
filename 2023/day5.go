package main

import (
	"aoc/util"
	"bufio"
	"golang.org/x/exp/slices"
	"io"
	"strconv"
	"strings"
)

func Day5(r io.Reader) (int, int, error) {
	type almanacCategoryEntry struct {
		sourceStart      int
		destinationStart int
		rangeLength      int
		sourceRange      [2]int
	}

	scanner := bufio.NewScanner(r)

	var seeds []int
	var almanac [][]almanacCategoryEntry

	var scanningCategory string
	var categoryIdx = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seedsData := strings.TrimPrefix(line, "seeds:")
			for _, seedS := range strings.Fields(seedsData) {
				seed, err := strconv.ParseInt(seedS, 10, 64)
				if err != nil {
					return 0, 0, err
				}

				seeds = append(seeds, int(seed))
			}
		} else if strings.HasSuffix(line, "map:") {
			if scanningCategory != "" {
				categoryIdx++
			}
			scanningCategory = strings.TrimSuffix(line, "map:")
			almanac = append(almanac, []almanacCategoryEntry{})
		} else if scanningCategory != "" {
			parts := strings.Fields(line)
			entry := almanacCategoryEntry{
				sourceStart:      util.MustInt(parts[1]),
				destinationStart: util.MustInt(parts[0]),
				rangeLength:      util.MustInt(parts[2]),
			}
			entry.sourceRange = [2]int{entry.sourceStart, entry.sourceStart + entry.rangeLength - 1}
			almanac[categoryIdx] = append(almanac[categoryIdx], entry)
		}
	}

	var minLocation = -1
	for _, seed := range seeds {
		var source = seed

		for _, c := range almanac {
			for _, mapping := range c {
				if source >= mapping.sourceStart && source <= mapping.sourceStart+mapping.rangeLength {
					offset := source - mapping.sourceStart
					source = mapping.destinationStart + offset
					break
				}
			}
		}

		if minLocation == -1 {
			minLocation = source
		} else {
			minLocation = min(minLocation, source)
		}
	}

	var minLocation2 = -1
	for i := 0; i < len(seeds); i += 2 {
		sourceRanges := [][2]int{{seeds[i], seeds[i] + seeds[i+1] - 1}}

		for _, c := range almanac {
			var mappedRanges [][2]int

			for len(sourceRanges) > 0 {
				sourceRange := sourceRanges[0]
				sourceRanges = slices.Delete(sourceRanges, 0, 1)
				var mapped bool
				for _, mapping := range c {
					// Fully below, not mapped yet
					if sourceRange[0] < mapping.sourceStart && sourceRange[1] < mapping.sourceStart {
						continue
					}
					// Fully above, not mapped yet
					if sourceRange[0] > mapping.sourceRange[1] && sourceRange[1] > mapping.sourceRange[1] {
						continue
					}
					// Fully within
					if sourceRange[0] >= mapping.sourceStart && sourceRange[1] <= mapping.sourceRange[1] {
						startOffset := sourceRange[0] - mapping.sourceStart
						endOffset := sourceRange[1] - mapping.sourceStart
						mappedRanges = append(mappedRanges, [2]int{mapping.destinationStart + startOffset, mapping.destinationStart + endOffset})
						mapped = true
						break
					}

					// Fully overlapping both endpoints, partially mapped, add ends to source ranges
					if sourceRange[0] < mapping.sourceStart && sourceRange[1] > mapping.sourceRange[1] {
						sourceRanges = append(sourceRanges,
							[2]int{sourceRange[0], mapping.sourceStart - 1},
							[2]int{mapping.sourceStart + mapping.rangeLength, sourceRange[1]},
						)
						mappedRanges = append(mappedRanges, [2]int{mapping.destinationStart, mapping.destinationStart + mapping.rangeLength - 1})
						mapped = true
						break
					}

					// Overlapping mapping start
					if sourceRange[0] < mapping.sourceStart {
						sourceRanges = append(sourceRanges, [2]int{sourceRange[0], mapping.sourceStart - 1})
						endOffset := sourceRange[1] - mapping.sourceStart
						mappedRanges = append(mappedRanges, [2]int{mapping.destinationStart, mapping.destinationStart + endOffset})
						mapped = true
						break
					}
					// Overlapping mapping end
					if sourceRange[1] > mapping.sourceRange[1] {
						sourceRanges = append(sourceRanges, [2]int{mapping.sourceStart + mapping.rangeLength, sourceRange[1]})
						startOffset := sourceRange[0] - mapping.sourceStart
						mappedRanges = append(mappedRanges, [2]int{mapping.destinationStart + startOffset, mapping.destinationStart + mapping.rangeLength})
						mapped = true
						break
					}
				}
				if !mapped {
					mappedRanges = append(mappedRanges, sourceRange)
				}
			}
			sourceRanges = mappedRanges
		}

		for _, r := range sourceRanges {
			if minLocation2 == -1 {
				minLocation2 = r[0]
			} else {
				minLocation2 = min(minLocation2, r[0])
			}
		}
	}

	return minLocation, minLocation2, nil
}
