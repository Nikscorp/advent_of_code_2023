package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	times, distances := parseInput(input)
	res := 1

	for i := 0; i < len(times); i++ {
		waysToWin := 0

		for holdTime := 1; holdTime < times[i]; holdTime++ {
			speed := holdTime
			remainingTime := times[i] - holdTime
			distance := speed * remainingTime

			if distance > distances[i] {
				waysToWin++
			}
		}

		res *= waysToWin
	}

	return res
}

func part2(input string) any {
	input = strings.ReplaceAll(input, " ", "")
	return part1(input)
}

func toInt(s string) int {
	res, _ := strconv.Atoi(s)

	return res
}

func parseInput(input string) ([]int, []int) {
	splitted := strings.Split(input, "\n")
	colonIndex := strings.Index(splitted[0], ":")
	timesFields := strings.Fields(splitted[0][colonIndex+1:])
	colonIndex = strings.Index(splitted[1], ":")
	distancesFields := strings.Fields(splitted[1][colonIndex+1:])

	times := make([]int, len(timesFields))
	distances := make([]int, len(timesFields))
	for i := range timesFields {
		times[i] = toInt(timesFields[i])
		distances[i] = toInt(distancesFields[i])
	}

	return times, distances
}
