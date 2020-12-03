// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

import (
	"strings"
	"testing"
)

func TestCountTrees(t *testing.T) {
	fixture := strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)
	grid, rowLength := ParseInput(fixture)
	type args struct {
		grid      Grid
		rowLength int
		slope     Slope
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{grid, rowLength, Slope{1, 1}}, 2},
		{"test 2", args{grid, rowLength, Slope{1, 3}}, 7},
		{"test 3", args{grid, rowLength, Slope{1, 5}}, 3},
		{"test 4", args{grid, rowLength, Slope{1, 7}}, 4},
		{"test 5", args{grid, rowLength, Slope{2, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTrees(tt.args.grid, tt.args.rowLength, tt.args.slope); got != tt.want {
				t.Errorf("CountTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountTrees(b *testing.B) {
	fixture := strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)
	grid, rowLength := ParseInput(fixture)
	for i := 0; i < b.N; i++ {
		CountTrees(grid, rowLength, Slope{1, 3})
	}
}
