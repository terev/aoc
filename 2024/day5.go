package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"aoc/util"
)

type pageRule struct {
	before []string
}

func parseRule(ruleRaw string, rules map[string]*pageRule) {
	l, r, f := strings.Cut(ruleRaw, "|")
	if !f {
		return
	}

	rule, ok := rules[l]
	if !ok {
		rule = &pageRule{}
		rules[l] = rule
	}

	rule.before = append(rule.before, r)
}

func Day5(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	var tot, tot2 int
	rules := make(map[string]*pageRule)
	parseSection := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(rules) > 0 && len(line) == 0 {
			parseSection = 1
			continue
		}

		switch parseSection {
		case 0:
			parseRule(line, rules)
		case 1:
			pages := strings.Split(line, ",")
			pagesCopy := slices.Clone(pages)
			slices.SortFunc(pagesCopy, func(a, b string) int {
				if a == b {
					return 0
				}
				if ar, ok := rules[a]; ok && slices.Contains(ar.before, b) {
					return -1
				}
				if br, ok := rules[b]; ok && slices.Contains(br.before, a) {
					return 1
				}
				return 0
			})
			if slices.Equal(pages, pagesCopy) {
				tot += util.MustInt(pages[len(pages)/2])
			} else {
				tot2 += util.MustInt(pagesCopy[len(pagesCopy)/2])
			}
		}
	}

	fmt.Println(tot)
	fmt.Println(tot2)

	return nil
}
