package main

import (
	"aoc/util"
	"bufio"
	"bytes"
	"io"
	"slices"
	"strings"
)

func Day13(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	var image [][]byte
	var p1, p2 int
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			transposedImage := util.TransposeMatrix(image)
			horizontal, vertical := findReflections(image), findReflections(transposedImage)
			p1 += vertical + horizontal*100
			horizontal, vertical = findReflectionsWithSmudges(image), findReflectionsWithSmudges(transposedImage)
			p2 += vertical + horizontal*100

			image = nil
			continue
		}
		image = append(image, []byte(line))
	}
	return p1, p2, nil
}

func findReflections(image [][]byte) int {
	var behindReflectionPoint int
	var potentialReflectionPoints []int
	for i := 1; i < len(image); i++ {
		for j, candidatePoint := range potentialReflectionPoints {
			oppositePoint := candidatePoint - (i - candidatePoint) - 1

			if oppositePoint >= 0 && !bytes.Equal(image[oppositePoint], image[i]) {
				potentialReflectionPoints = slices.Delete(potentialReflectionPoints, j, j+1)
			}
		}
		if bytes.Equal(image[i-1], image[i]) {
			potentialReflectionPoints = append(potentialReflectionPoints, i)
		}
	}
	if len(potentialReflectionPoints) > 0 {
		for _, p := range potentialReflectionPoints {
			behindReflectionPoint += p
		}
	}

	return behindReflectionPoint
}

func smudgeEqual(a []byte, b []byte) (bool, bool) {
	if len(a) != len(b) {
		panic("MUST BE SAME LEN")
	}

	var smudged bool
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		if smudged {
			return false, false
		}
		smudged = true
	}

	return true, smudged
}

type potentialReflectionPoint struct {
	index   int
	smudged bool
}

func findReflectionsWithSmudges(image [][]byte) int {
	var behindReflectionPoint int
	var potentialReflectionPoints []potentialReflectionPoint
	for i := 1; i < len(image); i++ {
		for j, candidatePoint := range potentialReflectionPoints {
			oppositePoint := candidatePoint.index - (i - candidatePoint.index) - 1

			if oppositePoint >= 0 {
				equal, withSmudge := smudgeEqual(image[oppositePoint], image[i])
				if !equal {
					potentialReflectionPoints = slices.Delete(potentialReflectionPoints, j, j+1)
					continue
				}

				if !withSmudge {
					continue
				}
				// if already smudged disregard point because there can only be one
				if candidatePoint.smudged {
					potentialReflectionPoints = slices.Delete(potentialReflectionPoints, j, j+1)
					continue
				}
				candidatePoint.smudged = true
				potentialReflectionPoints[j] = candidatePoint
			}
		}
		equal, withSmudge := smudgeEqual(image[i-1], image[i])
		if equal {
			potentialReflectionPoints = append(potentialReflectionPoints, potentialReflectionPoint{
				index:   i,
				smudged: withSmudge,
			})
		}
	}
	if len(potentialReflectionPoints) > 0 {
		for _, p := range potentialReflectionPoints {
			if p.smudged {
				behindReflectionPoint += p.index
			}
		}
	}

	return behindReflectionPoint
}
