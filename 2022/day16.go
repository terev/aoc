package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	valvePattern = `Valve\s+(\w+)\s+has flow rate=(\d+); tunnels? leads? to valves?\s+(.+)`
)

var (
	valveRe *regexp.Regexp
)

func init() {
	valveRe = regexp.MustCompile(valvePattern)
}

type vertex[T any] struct {
	edges map[string]int
	val   T
}

type graph[T any] struct {
	vertices map[string]*vertex[T]
}

func (g *graph[T]) AddVertex(v string, val T) {
	if g.vertices == nil {
		g.vertices = make(map[string]*vertex[T])
	}

	if _, exists := g.vertices[v]; !exists {
		g.vertices[v] = &vertex[T]{
			edges: make(map[string]int),
			val:   val,
		}
	}
}

func (g *graph[T]) AddEdge(from, to string, weight int) {
	g.vertices[from].edges[to] = weight
}

type valve struct {
	flowRate int
}

type traversalNode struct {
	cost   int
	vertex string
}

func (g *graph[T]) MinimalPath(from, to string) int {
	var toVisit = []traversalNode{{0, from}}

	visited := map[string]struct{}{
		from: {},
	}

	for len(toVisit) > 0 {
		cur := toVisit[0]
		toVisit = slices.Delete(toVisit, 0, 1)

		for edge := range g.vertices[cur.vertex].edges {
			if _, isVisited := visited[edge]; isVisited {
				continue
			}

			if edge == to {
				return cur.cost + 1
			}

			toVisit = append(toVisit, traversalNode{
				cost:   cur.cost + 1,
				vertex: edge,
			})

			visited[edge] = struct{}{}
		}
	}

	return -1
}

type traversalPlan struct {
	visited  int
	pressure int
	timeLeft int
	curValve string
}

func maximumPressureReleased(tunnels *graph[valve], distCache map[string]map[string]int, valveIndexes map[string]int, startValve string, initialTimeLeft, initiallyVisited int) int {
	plans := []traversalPlan{
		{
			curValve: startValve,
			timeLeft: initialTimeLeft,
			visited:  initiallyVisited,
			pressure: 0,
		},
	}

	if tunnels.vertices[startValve].val.flowRate > 0 {
		plans = append(plans, traversalPlan{
			curValve: startValve,
			timeLeft: initialTimeLeft - 1,
			visited:  initiallyVisited | (1 << valveIndexes[startValve]),
			pressure: tunnels.vertices[startValve].val.flowRate * (initialTimeLeft - 1),
		})
	}

	maxPressure := math.MinInt
	for len(plans) > 0 {
		plan := plans[0]
		plans = slices.Delete(plans, 0, 1)

		for nextValve, timeCost := range distCache[plan.curValve] {
			valveBit := 1 << valveIndexes[nextValve]
			if plan.visited&valveBit == valveBit {
				continue
			}

			timeLeft := plan.timeLeft - timeCost - 1
			if timeLeft < 0 {
				continue
			}

			newPressure := plan.pressure + timeLeft*tunnels.vertices[nextValve].val.flowRate

			if newPressure > maxPressure {
				maxPressure = newPressure
			}
			if timeLeft == 0 {
				continue
			}

			plans = slices.Insert(plans, 0, traversalPlan{
				curValve: nextValve,
				timeLeft: timeLeft,
				visited:  plan.visited | valveBit,
				pressure: newPressure,
			})
		}
	}

	return maxPressure
}

func Day16() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day16.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	closedValvesWithPressure := []string{}
	closedValveIndexLookup := map[string]int{}
	tunnels := &graph[valve]{}

	for scanner.Scan() {
		t := scanner.Text()

		matches := valveRe.FindStringSubmatch(t)

		tunnels.AddVertex(matches[1], valve{
			flowRate: util.MustInt(matches[2]),
		})

		if tunnels.vertices[matches[1]].val.flowRate > 0 {
			closedValveIndexLookup[matches[1]] = len(closedValveIndexLookup)
			closedValvesWithPressure = append(closedValvesWithPressure, matches[1])
		}

		for _, edge := range strings.Split(matches[3], ",") {
			tunnels.AddEdge(matches[1], strings.TrimSpace(edge), 1)
		}
	}

	dists := map[string]map[string]int{}

	for fromValve := range tunnels.vertices {
		for _, toValve := range closedValvesWithPressure {
			if fromValve != toValve {
				minCost := tunnels.MinimalPath(fromValve, toValve)
				if dists[fromValve] == nil {
					dists[fromValve] = map[string]int{}
				}
				dists[fromValve][toValve] = minCost
			}
		}
	}

	fmt.Println(maximumPressureReleased(tunnels, dists, closedValveIndexLookup, "AA", 30, 0))

	var bestWith2 = math.MinInt

	var allVisitedState = 0
	for i := 0; i < len(closedValveIndexLookup); i++ {
		allVisitedState |= 1 << i
	}

	// give each participant half the valves to visit
	for i := 0; i < (allVisitedState+1)/2; i++ {
		pressure := maximumPressureReleased(tunnels, dists, closedValveIndexLookup, "AA", 26, i) +
			maximumPressureReleased(tunnels, dists, closedValveIndexLookup, "AA", 26, allVisitedState-i)

		if pressure > bestWith2 {
			bestWith2 = pressure
		}
	}

	fmt.Println(bestWith2)

	return nil
}
