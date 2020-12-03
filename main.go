package main

import (
	"fmt"
	"os"

	"github.com/nmueller-realdigital/aoc2020-day03/aoc"
)

func main() {
	file, err := os.Open("data/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	grid, rowLength := aoc.ParseInput(file)
	result := 1
	slopes := []aoc.Slope{{Down: 1, Right: 1}, {Down: 1, Right: 3}, {Down: 1, Right: 5}, {Down: 1, Right: 7}, {Down: 2, Right: 1}}
	for _, s := range slopes {
		treeCount := aoc.CountTrees(grid, rowLength, s)
		fmt.Println("Slope: ", s, " Tree Count: ", treeCount)
		result *= treeCount
	}
	fmt.Println("Product of tree count: ", result)
}
