package main

import (
	"bufio"
	"golang.org/x/exp/slices"
	"io"
	"strconv"
	"unicode"
)

func Day3(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	var window [][]rune
	var previousLineSymbols []int
	var sum, sum2 int
	for scanner.Scan() {
		var newLineSymbols = []int{}
		line := scanner.Text()
		for i, c := range line {
			if !unicode.IsDigit(c) && c != '.' {
				newLineSymbols = append(newLineSymbols, i)
			}
		}

		lineRunes := []rune(line)

		for _, x := range previousLineSymbols {
			var partNumbers []int

			if len(window) == 2 {
				// process possible numbers above symbol. need 2 previous lines to do this
				for i := -1; i <= 1; i++ {
					if unicode.IsDigit(window[1][x+i]) {
						// expand number
						offset := x + i
						var number = []rune{window[1][offset]}
						var startFound, endFound bool
						for j := 1; !startFound || !endFound; j++ {
							if !startFound {
								if offset-j < 0 || !unicode.IsDigit(window[1][offset-j]) {
									startFound = true
								} else {
									number = slices.Insert(number, 0, window[1][offset-j])
									window[1][offset-j] = '.'
								}
							}
							if !endFound {
								if offset+j >= len(window[1]) || !unicode.IsDigit(window[1][offset+j]) {
									endFound = true
								} else {
									number = append(number, window[1][offset+j])
									window[1][offset+j] = '.'
								}
							}
						}
						n, err := strconv.ParseInt(string(number), 10, 64)
						if err != nil {
							return 0, 0, err
						}
						sum += int(n)
						partNumbers = append(partNumbers, int(n))
					}
				}
			}
			// process possible numbers below symbol
			for i := -1; i <= 1; i++ {
				if unicode.IsDigit(lineRunes[x+i]) {
					// expand number
					offset := x + i
					var number = []rune{lineRunes[offset]}
					var startFound, endFound bool
					for j := 1; !startFound || !endFound; j++ {
						if !startFound {
							if offset-j < 0 || !unicode.IsDigit(lineRunes[offset-j]) {
								startFound = true
							} else {
								number = slices.Insert(number, 0, lineRunes[offset-j])
								lineRunes[offset-j] = '.'
							}
						}
						if !endFound {
							if offset+j >= len(lineRunes) || !unicode.IsDigit(lineRunes[offset+j]) {
								endFound = true
							} else {
								number = append(number, lineRunes[offset+j])
								lineRunes[offset+j] = '.'
							}
						}
					}
					n, err := strconv.ParseInt(string(number), 10, 64)
					if err != nil {
						return 0, 0, err
					}
					sum += int(n)
					partNumbers = append(partNumbers, int(n))
				}
			}

			// process possible numbers left of symbol
			{
				var number []rune

				for offset := x - 1; offset >= 0 && unicode.IsDigit(window[0][offset]); offset-- {
					number = slices.Insert(number, 0, window[0][offset])
					window[0][offset] = '.'
				}
				if len(number) > 0 {
					n, err := strconv.ParseInt(string(number), 10, 64)
					if err != nil {
						return 0, 0, err
					}
					sum += int(n)
					partNumbers = append(partNumbers, int(n))
				}
			}
			// process possible numbers right of symbol
			{
				var number []rune
				for offset := x + 1; offset < len(window[0]) && unicode.IsDigit(window[0][offset]); offset++ {
					number = append(number, window[0][offset])
					window[0][offset] = '.'
				}
				if len(number) > 0 {
					n, err := strconv.ParseInt(string(number), 10, 64)
					if err != nil {
						return 0, 0, err
					}
					sum += int(n)
					partNumbers = append(partNumbers, int(n))
				}
			}

			if window[0][x] == '*' && len(partNumbers) == 2 {
				sum2 += partNumbers[0] * partNumbers[1]
			}
		}
		if len(window) >= 2 {
			window = slices.Delete(window, len(window)-1, len(window))
		}
		window = slices.Insert(window, 0, lineRunes)
		previousLineSymbols = newLineSymbols
	}

	for _, x := range previousLineSymbols {
		var partNumbers []int
		// process possible numbers above symbol. need 2 previous lines to do this
		for i := -1; i <= 1; i++ {
			if unicode.IsDigit(window[1][x+i]) {
				// expand number
				offset := x + i
				var number = []rune{window[1][offset]}
				var startFound, endFound bool
				for j := 1; !startFound || !endFound; j++ {
					if !startFound {
						if offset-j < 0 || !unicode.IsDigit(window[1][offset-j]) {
							startFound = true
						} else {
							number = slices.Insert(number, 0, window[1][offset-j])
							window[1][offset-j] = '.'
						}
					}
					if !endFound {
						if offset+j >= len(window[1]) || !unicode.IsDigit(window[1][offset+j]) {
							endFound = true
						} else {
							number = append(number, window[1][offset+j])
							window[1][offset+j] = '.'
						}
					}
				}
				n, err := strconv.ParseInt(string(number), 10, 64)
				if err != nil {
					return 0, 0, err
				}
				sum += int(n)
				partNumbers = append(partNumbers, int(n))
			}
		}
		// process possible numbers left of symbol
		{
			var number []rune

			for offset := x - 1; offset >= 0 && unicode.IsDigit(window[0][offset]); offset-- {
				number = slices.Insert(number, 0, window[0][offset])
				window[0][offset] = '.'
			}
			if len(number) > 0 {
				n, err := strconv.ParseInt(string(number), 10, 64)
				if err != nil {
					return 0, 0, err
				}
				sum += int(n)
				partNumbers = append(partNumbers, int(n))
			}
		}
		// process possible numbers right of symbol
		{
			var number []rune
			for offset := x + 1; offset < len(window[0]) && unicode.IsDigit(window[0][offset]); offset++ {
				number = append(number, window[0][offset])
				window[0][offset] = '.'
			}
			if len(number) > 0 {
				n, err := strconv.ParseInt(string(number), 10, 64)
				if err != nil {
					return 0, 0, err
				}
				sum += int(n)
				partNumbers = append(partNumbers, int(n))
			}
		}

		if window[0][x] == '*' && len(partNumbers) == 2 {
			sum2 += partNumbers[0] * partNumbers[1]
		}
	}

	return sum, sum2, nil
}
