// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

import (
	"bufio"
	"io"
)

func ParseInput(r io.Reader) (Grid, int) {
	scanner := bufio.NewScanner(r)
	treeLookup := Grid{}
	var row, rowLength int
	for scanner.Scan() {
		if len(scanner.Text()) > rowLength {
			rowLength = len(scanner.Text())
		}
		treeLookup[row] = make(map[int]bool)
		for idx, r := range scanner.Text() {
			if r == '#' {
				treeLookup[row][idx] = true
			}
		}
		row++
	}
	return treeLookup, rowLength
}
