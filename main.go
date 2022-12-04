package main

import (
	"aoc/cmd"
	"log"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		log.Fatalln(err)
	}
}
