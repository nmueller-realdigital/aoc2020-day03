// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

type Grid map[int]map[int]bool

type Slope struct {
	Down, Right int
}

func CountTrees(grid Grid, rowLength int, slope Slope) (treeCount int) {
	var column int
	for row := 0; row <= len(grid); row += slope.Down {
		if isTree := grid[row][column]; isTree {
			treeCount++
		}
		column = (column + slope.Right) % rowLength
	}
	return treeCount
}
