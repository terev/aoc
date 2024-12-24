package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"slices"
	"strings"

	"aoc/util"
)

func Day19(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	var patterns []string
	var designsToValidate []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if patterns == nil {
			patterns = strings.Split(strings.ReplaceAll(line, " ", ""), ",")
			continue
		}

		designsToValidate = append(designsToValidate, line)
	}

	var patternTrie util.Trie
	for i := 0; i < len(patterns); i++ {
		patternTrie.Insert(patterns[i])
	}

	var possibleDesigns int
	var totalWaysToMakeDesigns int
	for i := 0; i < len(designsToValidate); i++ {
		waysToMakeDesign := canDesignBeMade(designsToValidate[i], patternTrie)
		if waysToMakeDesign > 0 {
			possibleDesigns++
			totalWaysToMakeDesigns += waysToMakeDesign
			fmt.Println(waysToMakeDesign, designsToValidate[i])
		}
	}

	fmt.Println("Possible Designs:", possibleDesigns)
	fmt.Println("Possible Configurations:", totalWaysToMakeDesigns)
	return nil
}

type prefixSearchStack struct {
	prefixes  []string
	startIdx  int
	prefixIdx int
}

func canDesignBeMade(design string, patternTrie util.Trie) int {
	stackPointer := 0
	searchStack := []prefixSearchStack{{
		prefixes: patternTrie.Prefixes(design),
	}}

	designLen := len(design)

	// [idx] -> count
	validatedPrefixes := map[int]int{}

	for stackPointer >= 0 {
		stack := &searchStack[stackPointer]
		if stack.prefixIdx >= len(stack.prefixes) {
			if stackPointer > 0 {
				// Propagate count to previous stack.
				validatedPrefixes[searchStack[stackPointer-1].startIdx] += validatedPrefixes[stack.startIdx]
			}
			searchStack = slices.Delete(searchStack, stackPointer, stackPointer+1)
			stackPointer--
			continue
		}

		nextPrefix := stack.prefixes[stack.prefixIdx]
		stack.prefixIdx++

		nextI := stack.startIdx + len(nextPrefix)

		if nextI == designLen {
			validatedPrefixes[stack.startIdx]++
			continue
		}

		if totalFromPrefix, ok := validatedPrefixes[nextI]; ok {
			validatedPrefixes[stack.startIdx] += totalFromPrefix
			continue
		}

		nextPrefixes := slices.DeleteFunc(patternTrie.Prefixes(design[nextI:]),
			func(s string) bool {
				return nextI+len(s) > designLen
			})

		if len(nextPrefixes) > 0 {
			searchStack = append(searchStack, prefixSearchStack{
				prefixes: nextPrefixes,
				startIdx: nextI,
			})
			stackPointer++
			continue
		}
	}

	if len(validatedPrefixes) == 0 {
		return 0
	}

	return slices.Max(slices.Collect(maps.Values(validatedPrefixes)))
}
