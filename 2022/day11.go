package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	itemsPattern          = `\s*Starting items:\s*(.+$)`
	operationPattern      = `\s*Operation:\s*new\s*=\s*(old|\d+)\s*([+*])\s*(old|\d+)\s*`
	testPattern           = `\s*Test:\s*divisible by (\d+)`
	trueConditionPattern  = `\s*If true:\s*throw to monkey (\d+)`
	falseConditionPattern = `\s*If false:\s*throw to monkey (\d+)`
)

var (
	itemsRe, operationsRe, testRe, trueConditionRe, falseConditionRe *regexp.Regexp
)

func init() {
	itemsRe = regexp.MustCompile(itemsPattern)
	operationsRe = regexp.MustCompile(operationPattern)
	testRe = regexp.MustCompile(testPattern)
	trueConditionRe = regexp.MustCompile(trueConditionPattern)
	falseConditionRe = regexp.MustCompile(falseConditionPattern)
}

type operation struct {
	operand  string
	operand2 string
	operator string
}

type monkey struct {
	items            []int
	inspectOperation operation
	divisibilityTest int
	trueAction       int
	falseAction      int
}

func (m *monkey) receiveItem(itemWorry int) {
	m.items = append(m.items, itemWorry)
}

func (m *monkey) inspectNextItem(withRelief bool) (int, int, bool) {
	if len(m.items) == 0 {
		return 0, 0, false
	}

	item := m.items[0]
	if len(m.items) > 1 {
		m.items = m.items[1:]
	} else {
		m.items = []int{}
	}

	op1 := resolveOperandValue(m.inspectOperation.operand, item)
	op2 := resolveOperandValue(m.inspectOperation.operand2, item)

	switch m.inspectOperation.operator {
	case "+":
		item = op1 + op2
	case "*":
		item = op1 * op2
	default:
		panic("Unsupported operand")
	}

	if withRelief {
		item /= 3
	}

	if item%m.divisibilityTest == 0 {
		return item, m.trueAction, true
	}

	return item, m.falseAction, true
}

func resolveOperandValue(operand string, old int) int {
	switch operand {
	case "old":
		return old
	default:
		return util.MustInt(operand)
	}
}

func Day11() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day11.txt"))
	if err != nil {
		return err
	}

	var monkeys []monkey
	newMonkey := monkey{}

	// Divisors are all prime this means LCM is the product of them.
	// We can use this as the bound of worry.
	var commonMultiple = 1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		for i := 0; i < 5; i++ {
			scanner.Scan()
			t := scanner.Text()

			switch i {
			case 0:
				items := itemsRe.FindStringSubmatch(t)
				for _, item := range strings.Split(items[1], ",") {
					newMonkey.items = append(newMonkey.items, util.MustInt(strings.TrimSpace(item)))
				}
			case 1:
				opParts := operationsRe.FindStringSubmatch(t)
				newMonkey.inspectOperation = operation{
					operand:  opParts[1],
					operator: opParts[2],
					operand2: opParts[3],
				}
			case 2:
				divisibleTest := testRe.FindStringSubmatch(t)
				newMonkey.divisibilityTest = util.MustInt(divisibleTest[1])
				commonMultiple *= newMonkey.divisibilityTest
			case 3:
				trueCondition := trueConditionRe.FindStringSubmatch(t)
				newMonkey.trueAction = util.MustInt(trueCondition[1])
			case 4:
				falseCondition := falseConditionRe.FindStringSubmatch(t)
				newMonkey.falseAction = util.MustInt(falseCondition[1])
			}
		}
		monkeys = append(monkeys, newMonkey)
		newMonkey = monkey{}
	}

	var monkeysP2 = make([]monkey, len(monkeys))
	copy(monkeysP2, monkeys)

	inspectCounts := make([]int, len(monkeys))

	for round := 0; round < 20; round++ {
		for m := 0; m < len(monkeys); m++ {
			item, throwTo, wasInspected := monkeys[m].inspectNextItem(true)
			for wasInspected {
				monkeys[throwTo].receiveItem(item)
				inspectCounts[m]++
				item, throwTo, wasInspected = monkeys[m].inspectNextItem(true)
			}
		}
	}

	sort.Ints(inspectCounts)
	fmt.Println(inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2])

	inspectCounts = make([]int, len(monkeys))

	for round := 0; round < 10000; round++ {
		for m := 0; m < len(monkeysP2); m++ {
			item, throwTo, wasInspected := monkeysP2[m].inspectNextItem(false)
			for wasInspected {
				item %= commonMultiple
				monkeysP2[throwTo].receiveItem(item)
				inspectCounts[m]++
				item, throwTo, wasInspected = monkeysP2[m].inspectNextItem(false)
			}
		}
	}

	sort.Ints(inspectCounts)
	fmt.Println(inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2])

	return nil
}
