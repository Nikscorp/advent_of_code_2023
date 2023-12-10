package main

import (
	"strings"
)

type coord struct {
	i, j int
}

type dir coord

var (
	dirNorth = dir{-1, 0}
	dirSouth = dir{1, 0}
	dirWest  = dir{0, -1}
	dirEast  = dir{0, 1}
)

var pipes = map[byte]map[dir]dir{
	'|': {
		dirNorth: dirNorth,
		dirSouth: dirSouth,
	},
	'-': {
		dirEast: dirEast,
		dirWest: dirWest,
	},
	'L': {
		dirSouth: dirEast,
		dirWest:  dirNorth,
	},
	'J': {
		dirEast:  dirNorth,
		dirSouth: dirWest,
	},
	'7': {
		dirEast:  dirSouth,
		dirNorth: dirWest,
	},
	'F': {
		dirNorth: dirEast,
		dirWest:  dirSouth,
	},
}

func part1(input string) any {
	grid := parseInput(input)
	s := findStart(grid)
	loop := findLoop(s, grid)

	return len(loop) / 2
}

func part2(input string) any {
	grid := parseInput(input)
	s := findStart(grid)
	loop := findLoop(s, grid)

	// https://en.wikipedia.org/wiki/Shoelace_formula
	polygonArea := 0
	for i := 0; i < len(loop); i++ {
		cur := loop[i]
		next := loop[(i+1)%len(loop)]

		polygonArea += cur.i*next.j - cur.j*next.i
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	return polygonArea - len(loop)/2 + 1
}

func parseInput(input string) [][]byte {
	splitted := strings.Split(input, "\n")
	grid := make([][]byte, len(splitted))
	for i := range splitted {
		grid[i] = []byte(splitted[i])
	}

	return grid
}

func findStart(grid [][]byte) coord {
	var s coord

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				s.i = i
				s.j = j
				return s
			}
		}
	}

	panic("no S found in grid")
}

func findLoop(s coord, grid [][]byte) []coord {
	for _, pipe := range "|-LJ7F" {
		grid[s.i][s.j] = byte(pipe)
		loop := checkLoop(s, grid)
		if loop != nil {
			return loop
		}
	}

	panic("no loop found")
}

func checkLoop(s coord, grid [][]byte) []coord {
	cur := s
	dir := anyKey(pipes[grid[s.i][s.j]])

	res := []coord{}

	for {
		res = append(res, cur)
		newDir, ok := pipes[grid[cur.i][cur.j]][dir]
		if !ok {
			return nil
		}

		newCoord := coord{cur.i + newDir.i, cur.j + newDir.j}

		if newCoord.i < 0 || newCoord.i >= len(grid) || newCoord.j < 0 || newCoord.j >= len(grid[newCoord.i]) {
			return nil
		}
		if newCoord == s {
			if _, ok := pipes[grid[s.i][s.j]][newDir]; !ok {
				return nil
			}
			break
		}
		cur = newCoord
		dir = newDir
	}

	return res
}

func anyKey(m map[dir]dir) dir {
	for k := range m {
		return k
	}

	panic("empty map")
}
