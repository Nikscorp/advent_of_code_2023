package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0

	for _, line := range strings.Split(input, "\n") {
		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+2:]
		wantNumbersS, gotNumbersS, _ := strings.Cut(line, " | ")
		wantNumbers := strings.Split(wantNumbersS, " ")
		gotNumbers := strings.Split(gotNumbersS, " ")

		want := make(map[int]struct{}, len(wantNumbers))
		winCnt := 0

		for _, wantN := range wantNumbers {
			if wantN == "" {
				continue
			}
			want[toInt(wantN)] = struct{}{}
		}

		for _, gotN := range gotNumbers {
			if gotN == "" {
				continue
			}
			if _, ok := want[toInt(gotN)]; ok {
				winCnt++
			}
		}

		if winCnt > 0 {
			res += 1 << (winCnt - 1)
		}
	}

	return res
}

func part2(input string) any {
	res := 0
	additionalCards := make(map[int]int)

	for _, line := range strings.Split(input, "\n") {
		colonIndex := strings.Index(line, ":")
		_, gameIDS, _ := strings.Cut(line[:colonIndex], " ")
		gameID := toInt(strings.TrimSpace(gameIDS))

		res += 1 + additionalCards[gameID]

		wantNumbersS, gotNumbersS, _ := strings.Cut(line, " | ")
		wantNumbers := strings.Split(wantNumbersS, " ")
		gotNumbers := strings.Split(gotNumbersS, " ")

		want := make(map[int]struct{}, len(wantNumbers))
		winCnt := 0

		for _, wantN := range wantNumbers {
			if wantN == "" {
				continue
			}
			want[toInt(wantN)] = struct{}{}
		}

		for _, gotN := range gotNumbers {
			if gotN == "" {
				continue
			}
			if _, ok := want[toInt(gotN)]; ok {
				winCnt++
			}
		}

		winCopies := 1 + additionalCards[gameID]
		for i := 1; i <= winCnt; i++ {
			additionalCards[gameID+i] += winCopies
		}
	}

	return res
}

func toInt(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}
