package main

import (
	"strconv"
	"strings"
)

type coord struct {
	i, j int
}

var directions = map[byte]coord{
	'R': {0, 1},
	'L': {0, -1},
	'U': {-1, 0},
	'D': {1, 0},
}

func part1(input string) any {
	cur := coord{0, 0}
	coords := []coord{cur}

	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, " ")
		dir := directions[splitted[0][0]]
		n, _ := strconv.Atoi(splitted[1])
		for t := 0; t < n; t++ {
			cur.i += dir.i
			cur.j += dir.j

			coords = append(coords, cur)
		}
	}

	if coords[0] != coords[len(coords)-1] {
		panic("invalid input")
	}

	loop := coords[:len(coords)-1]
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

	return polygonArea + len(loop)/2 + 1
}

var directions2 = map[byte]coord{
	'0': {0, 1},
	'2': {0, -1},
	'3': {-1, 0},
	'1': {1, 0},
}

func part2(input string) any {
	cur := coord{0, 0}
	coords := []coord{cur}

	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, " ")
		dir := directions2[splitted[2][len(splitted[2])-2]]
		n, _ := strconv.ParseInt(splitted[2][2:len(splitted[2])-2], 16, 64)
		for t := 0; t < int(n); t++ {
			cur.i += dir.i
			cur.j += dir.j

			coords = append(coords, cur)
		}
	}

	if coords[0] != coords[len(coords)-1] {
		panic("invalid input")
	}

	loop := coords[:len(coords)-1]
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

	return polygonArea + len(loop)/2 + 1
}
