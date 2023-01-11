package main

import "testing"

func Test_towerHeight(t *testing.T) {
	towerHeight(100000, []byte(">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"))
}
