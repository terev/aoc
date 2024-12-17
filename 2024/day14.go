package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"aoc/util"
)

var robotRegex = regexp.MustCompile(`p=(\d+),(\d+)\s*v=(-?\d+),(-?\d+)`)

type bathroomRobot struct {
	pos      [2]int
	velocity [2]int
}

func Day14(in io.Reader, areaWidth, areaHeight, seconds int) error {
	robots := parsesRobotDescriptors(in)
	factor, _ := safetyFactorForSecond(robots, areaWidth, areaHeight, seconds)
	fmt.Println(factor)

	return nil
}

func parsesRobotDescriptors(in io.Reader) []bathroomRobot {
	scanner := bufio.NewScanner(in)

	var robots []bathroomRobot
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		botDescriptor := robotRegex.FindStringSubmatch(line)
		robots = append(robots, bathroomRobot{
			pos:      [2]int{util.MustInt(botDescriptor[1]), util.MustInt(botDescriptor[2])},
			velocity: [2]int{util.MustInt(botDescriptor[3]), util.MustInt(botDescriptor[4])},
		})
	}
	return robots
}

func safetyFactorForSecond(robots []bathroomRobot, areaWidth, areaHeight, second int) (int, []bathroomRobot) {
	midX := areaWidth / 2
	midY := areaHeight / 2

	var quadrants [4]int
	for i := 0; i < len(robots); i++ {
		robots[i].pos[0] = util.CircularIndex(robots[i].pos[0]+robots[i].velocity[0]*second, areaWidth)
		robots[i].pos[1] = util.CircularIndex(robots[i].pos[1]+robots[i].velocity[1]*second, areaHeight)
		if robots[i].pos[0] == midX || robots[i].pos[1] == midY {
			continue
		}

		x := util.BoolToByte(robots[i].pos[0] > midX)
		y := util.BoolToByte(robots[i].pos[1] > midY)
		// 1 1 - bottom right
		// 1 0 - bottom left
		// 0 1 - top right
		// 0 0 - top left
		quadrants[y<<1+x]++
	}

	factor := quadrants[0]
	for i := 1; i < len(quadrants); i++ {
		factor *= quadrants[i]
	}
	return factor, robots
}

type robotResult struct {
	second int
	factor int
	robots []bathroomRobot
}

func main() {
	areaWidth, areaHeight := 101, 103

	f, err := os.Open(filepath.Join(util.Cwd(), "day14.txt"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	robots := parsesRobotDescriptors(f)

	fmt.Println(robots[0])

	var results []robotResult
	for i := 0; i < 10000; i++ {
		factor, simulatedBots := safetyFactorForSecond(slices.Clone(robots), areaWidth, areaHeight, i)
		results = append(results, robotResult{second: i, factor: factor, robots: simulatedBots})
	}

	slices.SortFunc(results, func(a, b robotResult) int {
		return a.factor - b.factor
	})

	for s := 0; s < len(results); s++ {
		for i := 0; i < areaHeight; i++ {
			for j := 0; j < areaWidth; j++ {
				if slices.ContainsFunc(results[s].robots, func(robot bathroomRobot) bool {
					return robot.pos[0] == j && robot.pos[1] == i
				}) {
					fmt.Print("O")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println(results[s].factor, results[s].second)

		_, err := fmt.Scanln()
		if err != nil {
			panic(err)
		}
	}
}
