package main

import (
	"fmt"
	"io"
	"slices"

	"aoc/util"
)

func Day9(in io.Reader) error {
	l, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	// start, size
	var freeRegions [][2]int
	// start, size
	var files [][2]int

	fid := 0
	position := 0
	var fs []int

	for i := 0; i < len(l); i++ {
		size := util.MustInt(string(l[i]))
		if i%2 == 0 {
			fs = append(fs, slices.Repeat([]int{fid}, size)...)
			files = append(files, [2]int{position, size})
			fid++
		} else {
			fs = append(fs, slices.Repeat([]int{-1}, size)...)
			freeRegions = append(freeRegions, [2]int{position, size})
		}
		position += size
	}

	fs2 := slices.Clone(fs)

	writePointer := 0
	readPointer := position - 1
	for readPointer > writePointer {
		if fs[writePointer] != -1 {
			writePointer = writePointer + 1
			continue
		}

		if fs[readPointer] == -1 {
			readPointer--
			continue
		}

		fs[writePointer] = fs[readPointer]
		fs[readPointer] = -1
		writePointer = writePointer + 1
		readPointer--
	}

	var sum int
	for i := 0; i < len(fs); i++ {
		if fs[i] > 0 {
			sum += fs[i] * i
		}
	}

	fmt.Println(sum)

	for f := len(files) - 1; f > 0; f-- {
		for j := 0; j < len(freeRegions); j++ {
			if freeRegions[j][1] >= files[f][1] && freeRegions[j][0] < files[f][0] {
				for k := 0; k < files[f][1]; k++ {
					fs2[freeRegions[j][0]+k] = f
					fs2[files[f][0]+k] = -1
				}
				freeRegions[j][0] += files[f][1]
				freeRegions[j][1] -= files[f][1]
				break
			}
		}
	}

	var sum2 int
	for i := 0; i < len(fs2); i++ {
		if fs2[i] > 0 {
			sum2 += fs2[i] * i
		}
	}

	fmt.Println(sum2)
	return nil
}
