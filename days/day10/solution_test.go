package main

import "testing"

// FIXME
const exampleInput1_1 string = `.....
.S-7.
.|.|.
.L-J.
.....`

const exampleInput1_2 string = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

const exampleInput1_3 string = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const exampleInput1_4 string = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name:  "example",
			input: exampleInput1_1,
			want:  4,
		},
		{
			name:  "example",
			input: exampleInput1_2,
			want:  4,
		},
		{
			name:  "example",
			input: exampleInput1_3,
			want:  8,
		},
		{
			name:  "example",
			input: exampleInput1_4,
			want:  8,
		},
		{
			name:  "real",
			input: input,
			want:  6842,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

const exampleInput2_1 string = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const exampleInput2_2 string = `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`

const exampleInput2_3 string = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

const exampleInput2_4 string = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name:  "example",
			input: exampleInput2_1,
			want:  4,
		},
		{
			name:  "example",
			input: exampleInput2_2,
			want:  4,
		},
		{
			name:  "example",
			input: exampleInput2_3,
			want:  8,
		},
		{
			name:  "example",
			input: exampleInput2_4,
			want:  10,
		},
		{
			name:  "real",
			input: input,
			want:  393,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
