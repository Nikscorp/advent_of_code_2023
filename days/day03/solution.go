package main

import (
	"strings"
)

type coord struct {
	i, j int
}

var dirs = []coord{
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func part1(input string) any {
	matrix := strings.Split(input, "\n")
	res := 0

	for i := 0; i < len(matrix); i++ {
		curNum := 0
		isAdjacent := false
		for j := 0; j < len(matrix[i]); j++ {
			if !isDigit(matrix[i][j]) {
				if isAdjacent {
					res += curNum
				}
				curNum = 0
				isAdjacent = false
				continue
			}

			curNum = curNum*10 + int(matrix[i][j]-'0')
			if isAdjacent {
				continue
			}
			for _, dir := range dirs {
				newC := coord{i + dir.i, j + dir.j}
				if newC.i < 0 || newC.i >= len(matrix) || newC.j < 0 || newC.j >= len(matrix[newC.i]) {
					continue
				}

				curSym := matrix[newC.i][newC.j]
				if curSym == '.' || isDigit(curSym) {
					continue
				}
				isAdjacent = true
				break
			}
		}

		last := matrix[i][len(matrix[i])-1]
		if isDigit(last) && isAdjacent {
			res += curNum
		}
	}

	return res
}

type uniqueNum struct {
	num int
	i   int
	jTo int
}

func part2(input string) any {
	matrix := strings.Split(input, "\n")
	res := 0
	gearToNums := make(map[coord]map[uniqueNum]struct{})

	for i := 0; i < len(matrix); i++ {
		curNum := 0
		var adjGears []coord

		for j := 0; j < len(matrix[i]); j++ {
			if !isDigit(matrix[i][j]) {
				for _, gear := range adjGears {
					if gearToNums[gear] == nil {
						gearToNums[gear] = make(map[uniqueNum]struct{})
					}
					gearToNums[gear][uniqueNum{
						num: curNum,
						i:   i,
						jTo: j - 1,
					}] = struct{}{}
				}

				curNum = 0
				adjGears = adjGears[:0]

				continue
			}

			curNum = curNum*10 + int(matrix[i][j]-'0')

			for _, dir := range dirs {
				newC := coord{i + dir.i, j + dir.j}
				if newC.i < 0 || newC.i >= len(matrix) || newC.j < 0 || newC.j >= len(matrix[newC.i]) {
					continue
				}

				curSym := matrix[newC.i][newC.j]
				if curSym == '*' {
					adjGears = append(adjGears, newC)
				}
			}
		}

		last := matrix[i][len(matrix[i])-1]
		if isDigit(last) {
			for _, gear := range adjGears {
				if gearToNums[gear] == nil {
					gearToNums[gear] = make(map[uniqueNum]struct{})
				}
				gearToNums[gear][uniqueNum{
					num: curNum,
					i:   i,
					jTo: len(matrix[i]) - 1,
				}] = struct{}{}
			}
		}
	}

	for _, nums := range gearToNums {
		if len(nums) != 2 {
			continue
		}

		product := 1
		for n := range nums {
			product *= n.num
		}
		res += product
	}

	return res
}
