package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0

	for _, line := range strings.Split(input, "\n") {
		var (
			gameID int
			game   string
		)

		colonIndex := strings.IndexByte(line, ':')
		_, gameIDRow, _ := strings.Cut(line[:colonIndex], " ")
		gameID, _ = strconv.Atoi(gameIDRow)

		game = line[colonIndex+2:]

		rounds := strings.Split(game, "; ")
		ok := true
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			m := make(map[string]int, 3)

			for _, cube := range cubes {
				var (
					cnt   int
					cntS  string
					color string
				)

				cntS, color, _ = strings.Cut(cube, " ")
				cnt, _ = strconv.Atoi(cntS)

				m[color] += cnt
			}
			if m["red"] > 12 || m["green"] > 13 || m["blue"] > 14 {
				ok = false
				break
			}
		}

		if ok {
			res += gameID
		}
	}

	return res
}

func part2(input string) any {
	res := 0

	for _, line := range strings.Split(input, "\n") {
		game := line[strings.IndexByte(line, ':')+2:]

		rounds := strings.Split(game, "; ")

		var maxr, maxg, maxb int

		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			m := make(map[string]int, 3)
			for _, cube := range cubes {
				var (
					cnt   int
					cntS  string
					color string
				)

				cntS, color, _ = strings.Cut(cube, " ")
				cnt, _ = strconv.Atoi(cntS)
				m[color] += cnt
			}

			maxr = max(maxr, m["red"])
			maxg = max(maxg, m["green"])
			maxb = max(maxb, m["blue"])
		}

		res += maxr * maxg * maxb
	}

	return res
}
