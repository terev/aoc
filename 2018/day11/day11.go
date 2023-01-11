package main

import (
	"aoc/input"
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"os"
)

type Solution struct{}

func (s *Solution) Date() (int, int) {
	return 2018, 11
}

func (s *Solution) Execute(input Input) error {
	var serialNumber = cast.ToInt(string(bytes.TrimSpace(input.data)))

	var powerLevels [300][300]int

	for i := 0; i < len(powerLevels); i++ {
		for j := 0; j < len(powerLevels[i]); j++ {
			powerLevels[i][j] = powerLevel(j, i, serialNumber)
		}
	}

	var max = 0
	var maxPos = [3]int{0,0,0}

	var rows = len(powerLevels)
	var cols = len(powerLevels[0])

	for i := 0; i < cols; i++ {
		var temp = make([]int, rows)
		for j := i; j < cols; j++ {
			for k := 0; k < rows; k++ {
				temp[k] += powerLevels[k][j]
			}

			var localResult = kadane(temp)

			if localResult[0] > max {
				max = localResult[0]
				maxPos[0] = i
				maxPos[1] = localResult[1]
				maxPos[2] = (j - i) * (localResult[2] - localResult[1])
			}
		}
	}

	fmt.Println(max)
	fmt.Println(maxPos)

	return nil
}

func kadane(a []int) [3]int {
	if len(a) == 0 {
		return [3]int{0, 0, 0}
	}

	var result = [3]int{0, 0, 0}
	var localBegin = 0
	var localSum = 0

	for i := 0; i < len(a); i++ {
		localSum += a[i]

		if localSum <= result[0] {
			localSum = 0
			localBegin = i + 1
		} else {
			result[0] = localSum
			result[1] = localBegin
			result[2] = i
		}
	}

	if localSum > result[0] {
		result[0] = localSum
		result[1] = localBegin
		result[2] = len(a) - 1
	}

	return result
}

func powerLevel(x, y, serialNumber int) int {
	var rackID = x + 10
	return ((((rackID * y) + serialNumber) * rackID % 1000) / 100) - 5
}

var sample = `18`

type Input struct {
	data []byte
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	solution := Solution{}

	fmt.Println("Sample:")

	checkError(
		solution.Execute(Input{
			data: []byte(sample),
		}),
	)

	in, err := input.GetInput(solution.Date())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Real:")
	checkError(
		solution.Execute(Input{
			data: in,
		}),
	)
}
