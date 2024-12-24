package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"aoc/util"
)

func Day17(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	registers := map[string]int{}
	var program []int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		lh, rh, sep := strings.Cut(line, ":")
		if !sep {
			continue
		}
		if strings.HasPrefix(lh, "Register") {
			registers[strings.Fields(lh)[1]] = util.MustInt(strings.TrimSpace(rh))
		}
		if strings.HasPrefix(lh, "Program") {
			programRaw := strings.Split(strings.TrimSpace(rh), ",")
			for _, i := range programRaw {
				program = append(program, util.MustInt(i))
			}
		}
	}

	outputs := executeProgram(program, maps.Clone(registers))

	var outStr []string
	for _, o := range outputs {
		outStr = append(outStr, strconv.Itoa(o))
	}
	fmt.Println(strings.Join(outStr, ","))

	soln := 0
	loopmax := 8
	for i := len(program) - 1; i >= 0; i-- {
		var A int
		for j := 0; j < loopmax; j++ {
			A = soln + j
			reg2 := maps.Clone(registers)
			reg2["A"] = A
			if slices.Equal(program[i:], executeProgram(program, reg2)) {
				break
			}
		}

		loopmax *= 8
		soln = A * 8
	}

	fmt.Println("Part 2:", soln/8)

	return nil
}

func resolveComboOperand(operand int, registers map[string]int) int {
	if operand == 7 {
		panic("7 is an invalid operand")
	}

	if operand <= 3 {
		return operand
	}

	return registers[string(rune((operand-4)+'A'))]
}

func executeProgram(program []int, registers map[string]int) []int {
	var outputs []int

	instructionPtr := 0
	for instructionPtr < len(program) {
		opCode := program[instructionPtr]
		operand := program[instructionPtr+1]

		switch opCode {
		case 0:
			numerator := float64(registers["A"])
			denominator := math.Pow(2, float64(resolveComboOperand(operand, registers)))
			registers["A"] = int(numerator / denominator)
		case 1:
			registers["B"] ^= operand
		case 2:
			registers["B"] = resolveComboOperand(operand, registers) % 8
		case 3:
			if registers["A"] == 0 {
				break
			}
			instructionPtr = operand
			continue
		case 4:
			registers["B"] ^= registers["C"]
		case 5:
			outputs = append(outputs, resolveComboOperand(operand, registers)%8)
		case 6:
			numerator := float64(registers["A"])
			denominator := math.Pow(2, float64(resolveComboOperand(operand, registers)))
			registers["B"] = int(numerator / denominator)
		case 7:
			numerator := float64(registers["A"])
			denominator := math.Pow(2, float64(resolveComboOperand(operand, registers)))
			registers["C"] = int(numerator / denominator)
		}
		instructionPtr += 2
	}

	return outputs
}
