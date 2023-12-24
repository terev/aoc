package main

import (
	"bufio"
	"bytes"
	"golang.org/x/exp/slices"
	"io"
)

func Day14(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	var image [][]byte

	for scanner.Scan() {
		line := scanner.Bytes()
		if !bytes.Equal(line, []byte{}) {
			image = append(image, slices.Clone(line))
		}
	}

	endImage := simulateRocks(image)

	return calculateLoad(endImage), 0, nil
}

func simulateRocks(image [][]byte) [][]byte {
	var workingImage [][]byte
	for i := range image {
		workingImage = append(workingImage, slices.Clone(image[i]))
	}

	allSettled := false
	for !allSettled {
		allSettled = true
		for i := 1; i < len(workingImage); i++ {
			for j := 0; j < len(workingImage[i]); j++ {
				if workingImage[i][j] == 'O' && workingImage[i-1][j] == '.' {
					allSettled = false
					workingImage[i-1][j] = 'O'
					workingImage[i][j] = '.'
				}
			}
		}
	}

	return workingImage
}

func calculateLoad(image [][]byte) int {
	var load int
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == 'O' {
				load += len(image) - i
			}
		}
	}

	return load
}
