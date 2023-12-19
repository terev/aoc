package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type node struct {
	id    string
	left  *node
	right *node
}

func Day8(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	instructions := scanner.Text()

	network, err := parseNetwork(scanner)
	if err != nil {
		return 0, err
	}

	var stepCount int
	currentNode := network["AAA"]
	var instructionPointer int
	for currentNode.id != "ZZZ" {
		switch instructions[instructionPointer] {
		case 'L':
			currentNode = currentNode.left
		case 'R':
			currentNode = currentNode.right
		}

		instructionPointer = (instructionPointer + 1) % len(instructions)
		stepCount++
	}

	return stepCount, nil
}

func Day82(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	instructions := scanner.Text()
	network, err := parseNetwork(scanner)
	if err != nil {
		return 0, err
	}

	type mover struct {
		nodeStart   *node
		currentNode *node
		path        []*node
	}
	var movers []mover
	for k, n := range network {
		c := n
		if strings.HasSuffix(k, "A") {
			movers = append(movers, mover{nodeStart: c, currentNode: c})
		}
	}

	type cacheKey struct {
		fromNodeID    string
		instructionID int
	}
	type cacheValue struct {
		steps   int
		addedBy int
	}
	stepCountCache := make(map[cacheKey]cacheValue)

	var stepCount2 int
	var instructionPointer2 int
	var zCount int
	var endStepCount int
	var usedCache int

	nInstructions := len(instructions)

	for zCount != len(movers) {
		zCount = 0
		endStepCount = 0
		usedCache = 0

		for i := range movers {
			end := strings.HasSuffix(movers[i].currentNode.id, "Z")
			// Cache steps from each node+instruction
			if end {
				zCount++
				for j := len(movers[i].path) - 1; j >= 0; j-- {
					stepCountCache[cacheKey{
						fromNodeID:    movers[i].path[j].id,
						instructionID: (instructionPointer2 - (len(movers[i].path) - j - 1)) % nInstructions,
					}] = cacheValue{
						steps:   len(movers[i].path) - j,
						addedBy: i,
					}
				}
				movers[i].path = nil
			} else {
				if cache, ok := stepCountCache[cacheKey{
					fromNodeID:    movers[i].currentNode.id,
					instructionID: instructionPointer2,
				}]; ok {
					endStepCount += cache.steps
					zCount++
					usedCache++
				}
			}

			switch instructions[instructionPointer2] {
			case 'L':
				movers[i].path = append(movers[i].path, movers[i].currentNode.left)
				movers[i].currentNode = movers[i].currentNode.left
			case 'R':
				movers[i].path = append(movers[i].path, movers[i].currentNode.right)
				movers[i].currentNode = movers[i].currentNode.right
			}
		}

		instructionPointer2 = (instructionPointer2 + 1) % nInstructions
		stepCount2++
	}

	fmt.Println(stepCount2, endStepCount, stepCount2+endStepCount)

	return stepCount2 + endStepCount, nil
}

func parseNetwork(scanner *bufio.Scanner) (map[string]*node, error) {
	network := make(map[string]*node)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, "=")
		nodeID := strings.TrimSpace(parts[0])

		var currentNode *node
		if np, ok := network[nodeID]; ok {
			currentNode = np
		} else {
			currentNode = &node{
				id: nodeID,
			}
			network[nodeID] = currentNode
		}

		connections := strings.Split(strings.Trim(strings.TrimSpace(parts[1]), "()"), ",")
		leftID := strings.TrimSpace(connections[0])
		rightID := strings.TrimSpace(connections[1])

		if np, ok := network[leftID]; ok {
			currentNode.left = np
		} else {
			currentNode.left = &node{id: leftID}
			network[leftID] = currentNode.left
		}
		if np, ok := network[rightID]; ok {
			currentNode.right = np
		} else {
			currentNode.right = &node{id: rightID}
			network[rightID] = currentNode.right
		}
	}

	return network, nil
}
