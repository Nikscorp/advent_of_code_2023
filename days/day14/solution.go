package main

import (
	"strings"
)

func part1(input string) any {
	inputS := strings.Split(input, "\n")
	grid := make([][]byte, len(inputS))
	for i, line := range inputS {
		grid[i] = []byte(line)
	}

	return solve1(grid)
}

func part2(input string) any {
	inputS := strings.Split(input, "\n")
	grid := make([][]byte, len(inputS))
	for i, line := range inputS {
		grid[i] = []byte(line)
	}

	return solve2(grid)
}

func solve1(grid [][]byte) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] != 'O' {
				continue
			}

			fallI := i
			for fallI > 0 {
				if grid[fallI-1][j] != '.' {
					break
				}
				grid[fallI][j], grid[fallI-1][j] = grid[fallI-1][j], grid[fallI][j]
				fallI--
			}
		}
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'O' {
				res += (len(grid) - i)
			}
		}
	}

	return res
}

func solve2(grid [][]byte) int {
	memo := make(map[string]int)
	var gridS string

L:
	for i := 0; i < 1_000_000_000; i++ {
		cycle(grid)

		key := strings.Builder{}
		key.Grow(len(grid) * (len(grid[0]) + 1))
		for _, line := range grid {
			key.Write(line)
			key.WriteByte('\n')
		}

		keyS := key.String()

		if v, ok := memo[keyS]; ok {
			wantVV := (1_000_000_000-v)%(i-v) + v - 1
			for k, vv := range memo {
				if vv == wantVV {
					gridS = k
					break L
				}
			}
		}

		memo[keyS] = i
	}

	resGrid := strings.Split(gridS, "\n")
	resGrid = resGrid[:len(resGrid)-1]

	res := 0
	for i := 0; i < len(resGrid); i++ {
		for j := 0; j < len(resGrid[0]); j++ {
			if resGrid[i][j] == 'O' {
				res += (len(resGrid) - i)
			}
		}
	}

	return res
}

func cycle(grid [][]byte) {
	// north
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 'O' {
				continue
			}

			fallI := i
			for fallI > 0 {
				if grid[fallI-1][j] != '.' {
					break
				}
				grid[fallI][j], grid[fallI-1][j] = grid[fallI-1][j], grid[fallI][j]
				fallI--
			}
		}
	}

	// west
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != 'O' {
				continue
			}

			fallJ := j
			for fallJ > 0 {
				if grid[i][fallJ-1] != '.' {
					break
				}
				grid[i][fallJ], grid[i][fallJ-1] = grid[i][fallJ-1], grid[i][fallJ]
				fallJ--
			}
		}
	}

	// south
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 'O' {
				continue
			}

			fallI := i
			for fallI < len(grid)-1 {
				if grid[fallI+1][j] != '.' {
					break
				}
				grid[fallI][j], grid[fallI+1][j] = grid[fallI+1][j], grid[fallI][j]
				fallI++
			}
		}
	}

	// east
	for j := len(grid[0]) - 1; j >= 0; j-- {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != 'O' {
				continue
			}

			fallJ := j
			for fallJ < len(grid[0])-1 {
				if grid[i][fallJ+1] != '.' {
					break
				}
				grid[i][fallJ], grid[i][fallJ+1] = grid[i][fallJ+1], grid[i][fallJ]
				fallJ++
			}
		}
	}
}
