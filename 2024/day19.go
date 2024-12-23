package main

import (
	"bufio"
	"fmt"
	"io"
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

	slices.Sort(patterns)

	var patternTrie util.Trie
	for i := 0; i < len(patterns); i++ {
		patternTrie.Insert(patterns[i])
	}

	var possibleDesigns int
	for i := 0; i < len(designsToValidate); i++ {
		if canDesignBeMade(designsToValidate[i], patternTrie) {
			possibleDesigns++
			fmt.Println(designsToValidate[i])
		}
	}

	fmt.Println(possibleDesigns)
	return nil
}

func canDesignBeMade(design string, patternTrie util.Trie) bool {
	i := 0

	prefixStack := [][]string{patternTrie.Prefixes(design)}
	slices.Reverse(prefixStack[0])
	stackPointer := 0
	prefixPointers := []int{0}
	iStack := []int{i}

	for stackPointer >= 0 && iStack[stackPointer] < len(design) && len(prefixStack) > 0 && len(prefixStack[0]) > 0 {
		temp := stackPointer
		for stackPointer >= 0 && prefixPointers[stackPointer] >= len(prefixStack[stackPointer]) {
			prefixStack = slices.Delete(prefixStack, stackPointer, stackPointer+1)
			prefixPointers = slices.Delete(prefixPointers, stackPointer, stackPointer+1)
			iStack = slices.Delete(iStack, stackPointer, stackPointer+1)
			stackPointer--
		}
		if temp != stackPointer {
			continue
		}

		i = iStack[stackPointer]
		nextPrefix := prefixStack[stackPointer][prefixPointers[stackPointer]]
		prefixPointers[stackPointer]++
		if design[:i]+nextPrefix == design {
			return true
		}

		nextPrefixes := patternTrie.Prefixes(design[i+len(nextPrefix):])
		nextPrefixes = slices.DeleteFunc(nextPrefixes, func(s string) bool {
			return i+len(nextPrefix)+len(s) > len(design)
		})

		if len(nextPrefixes) > 0 {
			prefixStack = append(prefixStack, nextPrefixes)
			prefixPointers = append(prefixPointers, 0)
			iStack = append(iStack, i+len(nextPrefix))
			stackPointer++
		}
	}
	return i >= len(design)
}

func canMakeDesign2(design string, patternTrie util.Trie) bool {
	partsToMask := []string{design}
	for len(partsToMask) > 0 {
		var newParts []string
		for _, part := range partsToMask {
			longest := ""
			for i := 0; i < len(part); i++ {
				l := patternTrie.LongestPrefix(part[i:])
				if len(l) > len(longest) {
					longest = l
				}
			}
			if longest == "" {
				return false
			}
			for _, remainingPart := range strings.Split(part, longest) {
				if len(remainingPart) > 0 {
					newParts = append(newParts, remainingPart)
				}
			}
		}
		partsToMask = newParts
	}
	return true
}
