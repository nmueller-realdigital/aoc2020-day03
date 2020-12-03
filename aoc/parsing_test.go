// Package aoc implements business logic and utility functions for Advent of Code challenges.
package aoc

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
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
	got, got1 := ParseInput(fixture)
	want := make(map[int]map[int]bool)
	want[0] = map[int]bool{2: true, 3: true}
	want[1] = map[int]bool{0: true, 4: true, 8: true}
	want[2] = map[int]bool{1: true, 6: true, 9: true}
	want[3] = map[int]bool{2: true, 4: true, 8: true, 10: true}
	want[4] = map[int]bool{1: true, 5: true, 6: true, 9: true}
	want[5] = map[int]bool{2: true, 4: true, 5: true}
	want[6] = map[int]bool{1: true, 3: true, 5: true, 10: true}
	want[7] = map[int]bool{1: true, 10: true}
	want[8] = map[int]bool{0: true, 2: true, 3: true, 7: true}
	want[9] = map[int]bool{0: true, 4: true, 5: true, 10: true}
	want[10] = map[int]bool{1: true, 4: true, 8: true, 10: true}
	want1 := 11
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("ParseInput() got = %v, want %v", got, want)
	}
	if got1 != want1 {
		t.Errorf("ParseInput() got1 = %v, want %v", got1, want1)
	}
}
