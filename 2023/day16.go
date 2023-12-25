package main

import (
	"bufio"
	"bytes"
	"io"
	"slices"
)

func Day16(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	var contraption [][]byte

	for scanner.Scan() {
		line := scanner.Bytes()
		if !bytes.Equal(line, []byte{}) {
			contraption = append(contraption, slices.Clone(line))
		}
	}

	mostEnergized := 0
	for i := 0; i < len(contraption); i++ {
		mostEnergized = max(
			mostEnergized,
			simulateBeams(contraption, beam{position: [2]int{i, 0}, heading: 1}),
			simulateBeams(contraption, beam{position: [2]int{i, len(contraption[i]) - 1}, heading: 3}),
		)
	}

	for i := 0; i < len(contraption[0]); i++ {
		mostEnergized = max(
			mostEnergized,
			simulateBeams(contraption, beam{position: [2]int{0, i}, heading: 2}),
			simulateBeams(contraption, beam{position: [2]int{len(contraption) - 1, i}, heading: 0}),
		)
	}

	return simulateBeams(contraption, beam{position: [2]int{0, 0}, heading: 1}), mostEnergized, nil
}

type beam struct {
	position [2]int
	heading  uint8
}

type beamInterceptorTile interface {
	Bounce(b beam) []beam
}

type mirror struct {
	orientation byte
}

func (m mirror) Bounce(b beam) []beam {
	headingMap := map[byte][4]uint8{
		'/': {
			1,
			0,
			3,
			2,
		},
		'\\': {
			3,
			2,
			1,
			0,
		},
	}

	return []beam{
		{
			position: b.position,
			heading:  headingMap[m.orientation][b.heading],
		},
	}
}

type splitter struct {
	orientation byte
}

func (s splitter) Bounce(b beam) []beam {
	switch s.orientation {
	case '-':
		if b.heading == 1 || b.heading == 3 {
			return []beam{b}
		}
		return []beam{
			{position: b.position, heading: 1},
			{position: b.position, heading: 3},
		}
	case '|':
		if b.heading == 0 || b.heading == 2 {
			return []beam{b}
		}
		return []beam{
			{position: b.position, heading: 0},
			{position: b.position, heading: 2},
		}
	}
	panic("invalid orientation")
}

type air struct{}

func (air) Bounce(b beam) []beam {
	return []beam{b}
}

var tileTypes = map[byte]beamInterceptorTile{
	'/': mirror{
		orientation: '/',
	},
	'\\': mirror{
		orientation: '\\',
	},
	'-': splitter{
		orientation: '-',
	},
	'|': splitter{
		orientation: '|',
	},
	'.': air{},
}

func allBeamsInCache(beams []beam, cache map[beam]struct{}) bool {
	for _, curBeam := range beams {
		if _, ok := cache[curBeam]; !ok {
			return false
		}
	}

	return true
}

func simulateBeams(contraption [][]byte, initialBeam beam) int {
	var energizedMap = make([][]bool, len(contraption))
	for i := range contraption {
		energizedMap[i] = make([]bool, len(contraption[i]))
	}

	beamCache := make(map[beam]struct{})

	headingVectors := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	beams := []beam{initialBeam}

	var energizedTiles int
	for len(beams) > 0 && !allBeamsInCache(beams, beamCache) {
		var newBeams []beam
		for _, curBeam := range beams {
			if !energizedMap[curBeam.position[0]][curBeam.position[1]] {
				energizedTiles++
				energizedMap[curBeam.position[0]][curBeam.position[1]] = true
			}
			// If beam is in cache skip it because it's been handled
			if _, ok := beamCache[curBeam]; ok {
				continue
			}
			beamCache[curBeam] = struct{}{}

			bouncedBeams := tileTypes[contraption[curBeam.position[0]][curBeam.position[1]]].Bounce(curBeam)

			for _, newBeam := range bouncedBeams {
				curNewBeam := newBeam
				curNewBeam.position = [2]int{curNewBeam.position[0] + headingVectors[curNewBeam.heading][0], curNewBeam.position[1] + headingVectors[curNewBeam.heading][1]}
				if curNewBeam.position[0] < 0 || curNewBeam.position[0] >= len(contraption) ||
					curNewBeam.position[1] < 0 || curNewBeam.position[1] >= len(contraption[0]) {
					continue
				}
				newBeams = append(newBeams, curNewBeam)
			}
		}

		beams = slices.Clone(newBeams)
	}

	return energizedTiles
}
