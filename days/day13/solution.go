package main

import (
	"strings"
)

func part1(input string) any {
	res := 0
	for _, gridS := range strings.Split(input, "\n\n") {
		grid := strings.Split(gridS, "\n")

		res += solve1(grid)
	}

	return res
}

func part2(input string) any {
	res := 0
	for _, gridS := range strings.Split(input, "\n\n") {
		grid := strings.Split(gridS, "\n")

		res += solve2(grid)
	}

	return res
}

func solve1(grid []string) int {
	// vertical reflection
	for j := 0; j < len(grid[0])-1; j++ {
		left := j
		right := j + 1

		reflects := true
	L1:
		for left >= 0 && right < len(grid[0]) {
			for i := 0; i < len(grid); i++ {
				if grid[i][left] != grid[i][right] {
					reflects = false
					break L1
				}
			}
			left--
			right++
		}

		if reflects {
			return j + 1
		}
	}

	// horizontal reflection
	for i := 0; i < len(grid)-1; i++ {
		down := i
		up := i + 1

		reflects := true
	L2:
		for down >= 0 && up < len(grid) {
			for j := 0; j < len(grid[0]); j++ {
				if grid[down][j] != grid[up][j] {
					reflects = false
					break L2
				}
			}
			down--
			up++
		}

		if reflects {
			return (i + 1) * 100
		}
	}

	return 0
}

func solve2(grid []string) int {
	// vertical reflection
	for j := 0; j < len(grid[0])-1; j++ {
		left := j
		right := j + 1

		reflects := true
		smudgeCnt := 0
	L1:
		for left >= 0 && right < len(grid[0]) {
			for i := 0; i < len(grid); i++ {
				if grid[i][left] != grid[i][right] {
					if smudgeCnt == 0 {
						smudgeCnt++
					} else {
						reflects = false
						break L1
					}
				}
			}
			left--
			right++
		}

		if reflects && smudgeCnt == 1 {
			return j + 1
		}
	}

	// horizontal reflection
	for i := 0; i < len(grid)-1; i++ {
		down := i
		up := i + 1

		reflects := true
		smudgeCnt := 0
	L2:
		for down >= 0 && up < len(grid) {
			for j := 0; j < len(grid[0]); j++ {
				if grid[down][j] != grid[up][j] {
					if smudgeCnt == 0 {
						smudgeCnt++
					} else {
						reflects = false
						break L2
					}
				}
			}
			down--
			up++
		}

		if reflects && smudgeCnt == 1 {
			return (i + 1) * 100
		}
	}

	return 0
}
