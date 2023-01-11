package main

import (
	"aoc/input"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

type Solution struct{}

func (s *Solution) Date() (int, int) {
	return 2018, 7
}

type Step struct {
	id           string
	dependencies []*Step
	next         []*Step
	done         bool
	visits       int
}

func (s *Solution) Execute(input Input) error {
	scanner := bufio.NewScanner(bytes.NewReader(input.data))

	var steps = make(map[string]*Step)

	for scanner.Scan() {
		var stepID string
		var dependentStepID string
		_, err := fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.", &stepID, &dependentStepID)
		if err != nil {
			return err
		}

		var (
			step          *Step
			dependentStep *Step
		)

		if s, ok := steps[stepID]; ok {
			step = s
		} else {
			step = &Step{
				id:     stepID,
				visits: input.workRequiredFunc(stepID),
			}

			steps[stepID] = step
		}

		if s, ok := steps[dependentStepID]; ok {
			dependentStep = s
		} else {
			dependentStep = &Step{
				id:     dependentStepID,
				visits: input.workRequiredFunc(dependentStepID),
			}

			steps[dependentStepID] = dependentStep
		}

		step.next = append(step.next, dependentStep)
		sort.Slice(step.next, func(i, j int) bool {
			return step.next[i].id < step.next[j].id
		})

		dependentStep.dependencies = append(dependentStep.dependencies, step)
		sort.Slice(dependentStep.dependencies, func(i, j int) bool {
			return dependentStep.dependencies[i].id < dependentStep.dependencies[j].id
		})
	}

	var fringe []*Step

	for _, step := range steps {
		if len(step.dependencies) == 0 {
			fringe = append(fringe, step)
		}
	}

	sort.Slice(fringe, func(i, j int) bool {
		return fringe[i].id < fringe[j].id
	})

	var order string
	var total int
	var workers = make([]*Step, input.workerCount)
	for {
		var anyAssigned = false
		for _, worker := range workers {
			if worker != nil {
				anyAssigned = true
				break
			}
		}

		if len(fringe) == 0 && !anyAssigned {
			break
		}

		total++

		var next []*Step
		for i, worker := range workers {
			if worker == nil && len(fringe) > 0 {
				workers[i], fringe = fringe[0], fringe[1:]
				worker = workers[i]
			}

			if worker != nil {
				worker.visits--
				if worker.visits <= 0 {
					order += worker.id
					worker.done = true
					workers[i] = nil

					for _, step := range worker.next {
						var allDone = true
						for _, dep := range step.dependencies {
							if !dep.done {
								allDone = false
								break
							}
						}

						if allDone {
							next = append(next, step)
						}
					}
				}
			}
		}

		if len(next) > 0 {
			fringe = append(next, fringe...)

			sort.Slice(fringe, func(i, j int) bool {
				return fringe[i].id < fringe[j].id
			})
		}
	}

	fmt.Printf("Step order: %s, Total time: %d\n", order, total)

	return nil
}

type WorkRequiredFunc func(string) int

var sample = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

type Input struct {
	data             []byte
	workerCount      int
	workRequiredFunc WorkRequiredFunc
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
	fmt.Println("Part 1:")
	checkError(
		solution.Execute(Input{
			data:        []byte(sample),
			workerCount: 1,
			workRequiredFunc: func(stepID string) int {
				return 1
			},
		}),
	)

	fmt.Println("Part 2:")
	checkError(
		solution.Execute(Input{
			data:        []byte(sample),
			workerCount: 2,
			workRequiredFunc: func(stepID string) int {
				return int(stepID[0]) - 64
			},
		}),
	)

	in, err := input.GetInput(solution.Date())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Real:")
	fmt.Println("Part 1:")
	checkError(
		solution.Execute(Input{
			data:        in,
			workerCount: 1,
			workRequiredFunc: func(stepID string) int {
				return 1
			},
		}),
	)

	fmt.Println("Part 2:")
	checkError(
		solution.Execute(Input{
			data:        in,
			workerCount: 5,
			workRequiredFunc: func(stepID string) int {
				return int(stepID[0]) - 4
			},
		}),
	)
}
