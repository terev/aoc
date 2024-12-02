package main

import (
	"aoc/util"
	"bufio"
	"io"
	"strings"
)

func Day19(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	var parts []part
	workflows := make(map[string]workflow)
	scanState := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scanState = 1
			continue
		}

		if scanState == 0 {
			newWorkflow := parseWorkflow(line)
			workflows[newWorkflow.label] = newWorkflow
		} else {
			newPart := parsePart(line)
			parts = append(parts, newPart)
		}
	}

	var acceptedPartRatingSum int
	for _, curPart := range parts {
		tracker := partTracker{
			state:       PartStateActive,
			stateParams: "in",
			partRef:     &curPart,
		}

		for tracker.state == PartStateActive {
			curWorkflowLabel := tracker.stateParams
			for _, curRule := range workflows[tracker.stateParams].rules {
				curRule.Apply(&tracker)
				if tracker.stateParams != curWorkflowLabel || tracker.state != PartStateActive {
					break
				}
			}
		}

		if tracker.state == PartStateAccepted {
			for _, rating := range curPart.parameters {
				acceptedPartRatingSum += rating
			}
		}
	}

	return acceptedPartRatingSum, 0, nil
}

type unconditionalRule struct {
	effect effect
}

func (u unconditionalRule) Apply(p *partTracker) {
	u.effect.Apply(p)
}

type conditionalRule struct {
	leftOperand  string
	rightOperand string
	operator     byte
	effect       effect
}

func (c conditionalRule) Apply(p *partTracker) {
	if c.operator == '>' && p.partRef.parameters[c.leftOperand] > util.MustInt(c.rightOperand) {
		c.effect.Apply(p)
	} else if c.operator == '<' && p.partRef.parameters[c.leftOperand] < util.MustInt(c.rightOperand) {
		c.effect.Apply(p)
	}
}

type effectFinalRuling struct {
	accept bool
}

func (e effectFinalRuling) Apply(p *partTracker) {
	if e.accept {
		p.state = PartStateAccepted
	} else {
		p.state = PartStateRejected
	}
}

type effectJump struct {
	destination string
}

func (e effectJump) Apply(p *partTracker) {
	p.stateParams = e.destination
}

type partState int

const (
	PartStateActive partState = iota
	PartStateAccepted
	PartStateRejected
)

type partTracker struct {
	state       partState
	stateParams string
	partRef     *part
}

type effect interface {
	Apply(*partTracker)
}

type part struct {
	id         int
	parameters map[string]int
}

type rule interface {
	Apply(*partTracker)
}

type workflow struct {
	label string
	rules []rule
}

func parsePart(partData string) part {
	partData = strings.Trim(partData, "{}")
	partParameters := strings.Split(partData, ",")
	parameters := make(map[string]int)
	for _, parameterData := range partParameters {
		parameterParts := strings.Split(parameterData, "=")
		parameters[parameterParts[0]] = util.MustInt(parameterParts[1])
	}

	return part{
		parameters: parameters,
	}
}

func parseWorkflow(workflowData string) workflow {
	startBrace := strings.IndexByte(workflowData, '{')

	parsedWorkflow := workflow{label: workflowData[:startBrace]}
	content := workflowData[startBrace+1 : strings.IndexByte(workflowData, '}')]

	rulesData := strings.Split(content, ",")
	for i, ruleData := range rulesData {
		var newRule rule
		if i == len(rulesData)-1 {
			newRule = unconditionalRule{effect: parseEffect(ruleData)}
		} else {
			ruleParts := strings.Split(ruleData, ":")
			operatorIdx := strings.IndexAny(ruleParts[0], "<>")
			newRule = conditionalRule{
				effect:       parseEffect(ruleParts[1]),
				operator:     ruleParts[0][operatorIdx],
				leftOperand:  ruleParts[0][:operatorIdx],
				rightOperand: ruleParts[0][operatorIdx+1:],
			}
		}
		parsedWorkflow.rules = append(parsedWorkflow.rules, newRule)
	}

	return parsedWorkflow
}

func parseEffect(effectData string) effect {
	switch effectData {
	case "A":
		fallthrough
	case "R":
		return effectFinalRuling{accept: effectData == "A"}
	default:
		return effectJump{
			destination: effectData,
		}
	}
}
