package main

import (
	"aoc/input"
	//"aoc/input"
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"math"
	"os"
	"regexp"
	"sort"
)

type Solution struct{}

func (s *Solution) Date() (int, int) {
	return 2018, 10
}

type Vector struct {
	x int
	y int
}

type Velocity struct {
	xDelta int
	yDelta int
}

type Star struct {
	position Vector
	velocity Velocity
}

func (v *Vector) AddVelocity(vel Velocity) {
	v.x += vel.xDelta
	v.y += vel.yDelta
}

var starRegex = regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)>\s*velocity=<\s*(-?\d+),\s*(-?\d+)>`)

func (s *Solution) Execute(input Input) error {
	scanner := bufio.NewScanner(bytes.NewReader(input.data))

	var stars []*Star
	for scanner.Scan() {
		star := &Star{}
		parts := starRegex.FindStringSubmatch(scanner.Text())

		star.position = Vector{
			x: cast.ToInt(parts[1]),
			y: cast.ToInt(parts[2]),
		}

		star.velocity = Velocity{
			xDelta: cast.ToInt(parts[3]),
			yDelta: cast.ToInt(parts[4]),
		}

		stars = append(stars, star)
	}

	var second = 0
	for {
		if detectLetter(stars) {
			sky := plotStars(stars)

			for y := 0; y < len(sky); y++ {
				for x := 0; x < len(sky[y]); x++ {
					if sky[y][x] == 0 {
						fmt.Print(" ")
					} else {
						fmt.Print(string(sky[y][x]))
					}
				}
				fmt.Println()
			}
			fmt.Println(second)
			break
		}

		for _, star := range stars {
			star.position.AddVelocity(star.velocity)
		}

		second++
	}

	return nil
}

func detectLetter(stars []*Star) bool {
	// key is y coordinate, val is x coordinates
	var linesHorizontal = map[int][]int{}

	// key is x coordinate, val is y coordinates
	var linesVertical = map[int][]int{}

	for _, star := range stars {
		if _, ok := linesHorizontal[star.position.y]; ok {
			linesHorizontal[star.position.y] = append(linesHorizontal[star.position.y], star.position.x)
		} else {
			linesHorizontal[star.position.y] = []int{star.position.x}
		}

		if _, ok := linesVertical[star.position.x]; ok {
			linesVertical[star.position.x] = append(linesVertical[star.position.x], star.position.y)
		} else {
			linesVertical[star.position.x] = []int{star.position.y}
		}
	}
	//fmt.Println(linesHorizontal)
	//fmt.Println(linesVertical)

	var horizontalSegments [][3]int
	for y, xs := range linesHorizontal {
		if len(xs) <= 1 {
			continue
		}

		sort.Slice(xs, func(i, j int) bool {
			return xs[i] < xs[j]
		})

		prev := xs[0]
		first := prev
		for _, next := range xs {
			if next != prev && math.Abs(float64(next-prev)) > 1 {
				if first != prev && math.Abs(float64(prev-first)) > 3 {
					fmt.Println(prev-first)
					horizontalSegments = append(horizontalSegments, [3]int{y, first, prev})
				}
				first = next
			}

			prev = next
		}

		if first != prev && math.Abs(float64(prev-first)) > 3 {
			fmt.Println(prev-first)
			horizontalSegments = append(horizontalSegments, [3]int{y, first, prev})
		}
	}

	if len(horizontalSegments) < 1 {
		return false
	}

	var verticalSegments [][3]int
	for x, ys := range linesVertical {
		if len(ys) <= 1 {
			continue
		}

		sort.Slice(ys, func(i, j int) bool {
			return ys[i] < ys[j]
		})

		prev := ys[0]
		first := prev
		for _, next := range ys {
			if next != prev && math.Abs(float64(next-prev)) > 1 {
				if first != prev && math.Abs(float64(prev-first)) > 5 {
					fmt.Println(prev-first)
					verticalSegments = append(verticalSegments, [3]int{x, first, prev})
				}
				first = next
			}

			prev = next
		}

		if first != prev && math.Abs(float64(prev-first)) > 5 {
			fmt.Println(prev-first)
			verticalSegments = append(verticalSegments, [3]int{x, first, prev})
		}
	}

	if len(verticalSegments) < 1 {
		return false
	}

	fmt.Println(horizontalSegments)
	fmt.Println(verticalSegments)

	//for _, candidateh := range horizontalSegments {
	//	for _, candidatev := range verticalSegments {
	//		if candidatev[0] < candidateh[1] || candidatev[0] > candidateh[2] {
	//			continue
	//		}
	//
	//		if math.Abs(float64(candidatev[1]-candidateh[0])) <= 1 || math.Abs(float64(candidatev[2]-candidateh[0])) <= 1 {
	//			return true
	//		}
	//	}
	//}

	return true
}

func plotStars(stars []*Star) [][]byte {
	if len(stars) < 1 {
		return nil
	}

	var (
		minx = stars[0].position.x
		maxx = stars[0].position.x
		miny = stars[0].position.y
		maxy = stars[0].position.y
	)

	for i := 1; i < len(stars); i++ {
		if stars[i].position.x < minx {
			minx = stars[i].position.x
		}

		if stars[i].position.y < miny {
			miny = stars[i].position.y
		}

		if stars[i].position.x > maxx {
			maxx = stars[i].position.x
		}

		if stars[i].position.y > maxy {
			maxy = stars[i].position.y
		}
	}

	sky := make([][]byte, maxy-miny + 1)
	for i := 0; i < len(sky); i++ {
		sky[i] = make([]byte, maxx-minx + 1)
	}

	for i := 0; i < len(stars); i++ {
		sky[stars[i].position.y-miny][stars[i].position.x-minx] = '#'
	}

	return sky
}

//var sample = `position=< 0,  0> velocity=< 0,  0>
//position=< 1,  0> velocity=< 0,  0>
//position=< 2,  0> velocity=< 0,  0>
//position=< 1,  1> velocity=< 0,  0>
//position=< 1,  2> velocity=< 0,  0>
//position=< 0,  3> velocity=< 0,  0>
//position=< 1,  3> velocity=< 0,  0>
//position=< 2,  3> velocity=< 0,  0>`

var sample = `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>`

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
