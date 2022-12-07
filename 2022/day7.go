package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const typeFile = 1
const typeDir = 2

type node interface{}

type fileNode struct {
	name string
	size int
}

type dirNode struct {
	prevDir  *dirNode
	name     string
	children []node
}

func sumDirectorySizes(dir *dirNode, maxSize int, accum *int) int {
	if len(dir.children) == 0 {
		return 0
	}

	totalSize := 0

	for _, child := range dir.children {
		if f, ok := child.(*fileNode); ok {
			totalSize += f.size
		}

		if d, ok := child.(*dirNode); ok {
			totalSize += sumDirectorySizes(d, maxSize, accum)
		}
	}

	if totalSize <= maxSize {
		*accum += totalSize
	}

	return totalSize
}

func findMinRequiredDir(dir *dirNode, minSize int, min *int) int {
	if len(dir.children) == 0 {
		return 0
	}

	totalSize := 0

	for _, child := range dir.children {
		if f, ok := child.(*fileNode); ok {
			totalSize += f.size
		}

		if d, ok := child.(*dirNode); ok {
			totalSize += findMinRequiredDir(d, minSize, min)
		}
	}

	if totalSize >= minSize && totalSize < *min {
		*min = totalSize
	}

	return totalSize
}

func Day7() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day7.txt"))
	if err != nil {
		return err
	}

	root := &dirNode{name: "/"}
	curDir := root
	scanner := bufio.NewScanner(f)

scan:
	for scanner.Scan() {
		t := scanner.Text()
		tokens := strings.Split(t, " ")

		switch tokens[0] {
		case "$":
			switch tokens[1] {
			case "cd":
				switch tokens[2] {
				case "/":
					curDir = root
				case "..":
					curDir = curDir.prevDir
				default:
					for _, node := range curDir.children {
						if dir, ok := node.(*dirNode); ok && dir.name == tokens[2] {
							curDir = dir
							continue scan
						}
					}
					newNode := &dirNode{
						prevDir: curDir,
						name:    tokens[2],
					}
					curDir.children = append(curDir.children, newNode)
					curDir = newNode
				}
			case "ls":
			default:
				return fmt.Errorf("Invalid command %s", tokens[1])
			}
		case "dir":
			for _, node := range curDir.children {
				if dir, ok := node.(*dirNode); ok && dir.name == tokens[1] {
					continue scan
				}
			}
			curDir.children = append(curDir.children, &dirNode{
				prevDir: curDir,
				name:    tokens[1],
			})
		default:
			for _, node := range curDir.children {
				if f, ok := node.(*fileNode); ok && f.name == tokens[1] {
					continue scan
				}
			}

			fSize, err := strconv.ParseInt(tokens[0], 10, 64)
			if err != nil {
				return err
			}
			curDir.children = append(curDir.children, &fileNode{
				name: tokens[1],
				size: int(fSize),
			})
		}
	}

	var result int
	totalUsed := sumDirectorySizes(root, 100000, &result)
	fmt.Println(result)

	unused := 70000000 - totalUsed
	requiredSizeToDelete := 30000000 - unused

	min := 70000000
	findMinRequiredDir(root, requiredSizeToDelete, &min)

	fmt.Println(min)

	return nil
}
