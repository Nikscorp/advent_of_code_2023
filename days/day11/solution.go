package main

import (
	"math"
	"strings"
)

func part1(input string) any {
	grid := strings.Split(input, "\n")
	expandedRows := make([]bool, len(grid))
	expandedCols := make([]bool, len(grid[0]))
	galaxyInd := 0
	galaxies := make([][]int, len(grid))
	for i := range galaxies {
		galaxies[i] = make([]int, len(grid[0]))

		for j := range galaxies[i] {
			galaxies[i][j] = -1
		}
	}

	for i := 0; i < len(grid); i++ {
		expanded := true
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				expanded = false
				galaxies[i][j] = galaxyInd
				galaxyInd++
			}
		}

		expandedRows[i] = expanded
	}

	for j := 0; j < len(grid[0]); j++ {
		expanded := true
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != '.' {
				expanded = false
				break
			}
		}

		expandedCols[j] = expanded
	}

	res := 0

	for i := range galaxies {
		for j := range galaxies[i] {
			if galaxies[i][j] == -1 {
				continue
			}

			res += bfs(coord{i, j}, galaxies, galaxyInd, grid, expandedRows, expandedCols, 2-1)
			galaxies[i][j] = -1
		}
	}

	return res
}

func part2(input string) any {
	grid := strings.Split(input, "\n")
	expandedRows := make([]bool, len(grid))
	expandedCols := make([]bool, len(grid[0]))
	galaxyInd := 0
	galaxies := make([][]int, len(grid))
	for i := range galaxies {
		galaxies[i] = make([]int, len(grid[0]))

		for j := range galaxies[i] {
			galaxies[i][j] = -1
		}
	}

	for i := 0; i < len(grid); i++ {
		expanded := true
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				expanded = false
				galaxies[i][j] = galaxyInd
				galaxyInd++
			}
		}

		expandedRows[i] = expanded
	}

	for j := 0; j < len(grid[0]); j++ {
		expanded := true
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != '.' {
				expanded = false
				break
			}
		}

		expandedCols[j] = expanded
	}

	res := 0

	for i := range galaxies {
		for j := range galaxies[i] {
			if galaxies[i][j] == -1 {
				continue
			}

			res += bfs(coord{i, j}, galaxies, galaxyInd, grid, expandedRows, expandedCols, 1000000-1)
			galaxies[i][j] = -1
		}
	}

	return res
}

type coord struct {
	i, j int
}

type coordWithPath struct {
	coord
	curLen int
}

var dirs = []coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func bfs(start coord, ends [][]int, galaxiesCnt int, grid []string, expandedRows []bool, expandedCols []bool, delta int) int {
	var queue = []coordWithPath{{coord: start, curLen: 0}}

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	lens := make([]int, galaxiesCnt)
	for i := range lens {
		lens[i] = math.MaxInt
	}

	lens[ends[start.i][start.j]] = 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			newC := coordWithPath{coord: coord{cur.i + dir.i, cur.j + dir.j}, curLen: cur.curLen + 1}
			if newC.i < 0 || newC.i >= len(grid) || newC.j < 0 || newC.j >= len(grid[0]) {
				continue
			}

			if visited[newC.i][newC.j] {
				continue
			}

			if expandedRows[newC.i] {
				newC.curLen += delta
			}

			if expandedCols[newC.j] {
				newC.curLen += delta
			}

			if v := ends[newC.i][newC.j]; v != -1 {
				lens[v] = min(lens[v], newC.curLen)
			}

			visited[newC.i][newC.j] = true
			queue = append(queue, newC)
		}
	}

	res := 0
	for _, len := range lens {
		if len == math.MaxInt {
			continue
		}
		res += len
	}

	return res
}
