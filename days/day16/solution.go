package main

import (
	"math"
	"strings"
)

func part1(input string) any {
	grid := strings.Split(input, "\n")
	return bfs(grid, coordWithDir{coord{0, -1}, Right})
}

func part2(input string) any {
	res := math.MinInt

	grid := strings.Split(input, "\n")
	// leftmost or rightmost column
	for i := 0; i < len(grid); i++ {
		res = max(res, bfs(grid, coordWithDir{coord{i, -1}, Right}))
		res = max(res, bfs(grid, coordWithDir{coord{i, len(grid[0])}, Left}))
	}

	// top row or bottom row
	for j := 0; j < len(grid[0]); j++ {
		res = max(res, bfs(grid, coordWithDir{coord{-1, j}, Down}))
		res = max(res, bfs(grid, coordWithDir{coord{len(grid), j}, Up}))
	}

	return res
}

type coord struct {
	i, j int
}

type Direction coord

var (
	Any   = Direction{0, 0}
	Up    = Direction{-1, 0}
	Down  = Direction{1, 0}
	Left  = Direction{0, -1}
	Right = Direction{0, 1}
)

type coordWithDir struct {
	coord
	dir Direction
}

func bfs(grid []string, start coordWithDir) int {
	visited := make(map[coordWithDir]bool)
	visited[start] = true

	var queue = []coordWithDir{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curDir := cur.dir

		newC := coord{cur.i + curDir.i, cur.j + curDir.j}
		if newC.i < 0 || newC.i >= len(grid) || newC.j < 0 || newC.j >= len(grid[0]) {
			continue
		}

		if grid[newC.i][newC.j] == '.' {
			t := coordWithDir{newC, curDir}
			if visited[t] {
				continue
			}

			visited[t] = true
			queue = append(queue, t)
			continue
		}

		if grid[newC.i][newC.j] == '|' {
			if curDir == Up || curDir == Down {
				t := coordWithDir{newC, curDir}
				if visited[t] {
					continue
				}

				visited[t] = true
				queue = append(queue, t)
				continue
			} else {
				t1 := coordWithDir{newC, Up}
				if !visited[t1] {
					visited[t1] = true
					queue = append(queue, t1)
				}
				t2 := coordWithDir{newC, Down}
				if !visited[t2] {
					visited[t2] = true
					queue = append(queue, t2)
				}
			}
		}

		if grid[newC.i][newC.j] == '-' {
			if curDir == Left || curDir == Right {
				t := coordWithDir{newC, curDir}
				if visited[t] {
					continue
				}

				visited[t] = true
				queue = append(queue, t)
				continue
			} else {
				t1 := coordWithDir{newC, Left}
				if !visited[t1] {
					visited[t1] = true
					queue = append(queue, t1)
				}
				t2 := coordWithDir{newC, Right}
				if !visited[t2] {
					visited[t2] = true
					queue = append(queue, t2)
				}
			}
		}

		if grid[newC.i][newC.j] == '/' {
			newDir := Any
			switch curDir {
			case Up:
				newDir = Right
			case Down:
				newDir = Left
			case Left:
				newDir = Down
			case Right:
				newDir = Up
			}

			t := coordWithDir{newC, newDir}
			if visited[t] {
				continue
			}

			visited[t] = true
			queue = append(queue, t)
			continue
		}

		if grid[newC.i][newC.j] == '\\' {
			newDir := Any
			switch curDir {
			case Up:
				newDir = Left
			case Down:
				newDir = Right
			case Left:
				newDir = Up
			case Right:
				newDir = Down
			}

			t := coordWithDir{newC, newDir}
			if visited[t] {
				continue
			}

			visited[t] = true
			queue = append(queue, t)
			continue
		}
	}

	resMap := make(map[coord]bool)
	for v := range visited {
		resMap[v.coord] = true
	}

	return len(resMap) - 1
}
