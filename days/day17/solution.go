package main

import (
	"math"
	"strings"
)

func part1(input string) any {
	grid := strings.Split(input, "\n")

	return bfs1(grid)
}

func part2(input string) any {
	grid := strings.Split(input, "\n")

	return bfs2(grid)
}

type coord struct {
	i, j int
}

type bfsState struct {
	c          coord
	prevDir    Direction
	sameDirCnt int
	// heatLoss   int
}

type Direction coord

var (
	Any   = Direction{0, 0}
	Up    = Direction{-1, 0}
	Down  = Direction{1, 0}
	Left  = Direction{0, -1}
	Right = Direction{0, 1}
)

var opposite = map[Direction]Direction{
	Up:    Down,
	Down:  Up,
	Left:  Right,
	Right: Left,
	Any:   Any,
}

func bfs1(grid []string) int {
	start := bfsState{coord{0, 0}, Any, 100}

	queue := []bfsState{start}
	visited := make(map[bfsState]int)
	visited[start] = 0

	end := coord{len(grid) - 1, len(grid[0]) - 1}

	res := math.MaxInt

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.c == end {
			res = min(res, visited[cur])
			continue
		}

		for _, dir := range []Direction{Up, Down, Left, Right} {
			if cur.prevDir == dir && cur.sameDirCnt >= 3 {
				continue
			}

			if opposite[dir] == cur.prevDir {
				continue
			}

			newC := coord{cur.c.i + dir.i, cur.c.j + dir.j}
			if newC.i < 0 || newC.i >= len(grid) || newC.j < 0 || newC.j >= len(grid[0]) {
				continue
			}

			heatLoss := visited[cur] + int(grid[newC.i][newC.j]-'0')
			newState := bfsState{newC, dir, 1}
			if dir == cur.prevDir {
				newState.sameDirCnt = cur.sameDirCnt + 1
			}

			if v, ok := visited[newState]; ok && heatLoss >= v {
				continue
			}

			visited[newState] = heatLoss
			queue = append(queue, newState)
		}
	}

	return res
}

func bfs2(grid []string) int {
	start := bfsState{coord{0, 0}, Any, 100}

	queue := []bfsState{start}
	visited := make(map[bfsState]int)
	visited[start] = 0

	end := coord{len(grid) - 1, len(grid[0]) - 1}

	res := math.MaxInt

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.sameDirCnt < 4 {
			dir := cur.prevDir

			newC := coord{cur.c.i + dir.i, cur.c.j + dir.j}
			if newC.i < 0 || newC.i >= len(grid) || newC.j < 0 || newC.j >= len(grid[0]) {
				continue
			}
			newState := bfsState{newC, dir, cur.sameDirCnt + 1}
			heatLoss := visited[cur] + int(grid[newC.i][newC.j]-'0')
			if v, ok := visited[newState]; ok && heatLoss >= v {
				continue
			}

			visited[newState] = heatLoss
			queue = append(queue, newState)

			continue
		}

		if cur.c == end {
			res = min(res, visited[cur])
			continue
		}

		for _, dir := range []Direction{Up, Down, Left, Right} {
			if cur.prevDir == dir && cur.sameDirCnt >= 10 {
				continue
			}

			if opposite[dir] == cur.prevDir {
				continue
			}

			newC := coord{cur.c.i + dir.i, cur.c.j + dir.j}
			if newC.i < 0 || newC.i >= len(grid) || newC.j < 0 || newC.j >= len(grid[0]) {
				continue
			}

			heatLoss := visited[cur] + int(grid[newC.i][newC.j]-'0')
			newState := bfsState{newC, dir, 1}
			if dir == cur.prevDir {
				newState.sameDirCnt = cur.sameDirCnt + 1
			}

			if v, ok := visited[newState]; ok && heatLoss >= v {
				continue
			}

			visited[newState] = heatLoss
			queue = append(queue, newState)
		}
	}

	return res
}
