package main

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) any {
	times, distances := parseInput(input)
	res := 1

	for i := 0; i < len(times); i++ {
		waysToWinGot := solveEquation(distances[i], times[i])
		res *= waysToWinGot
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

func solveEquation(distance int, time int) int {
	d := math.Sqrt(float64(time*time - 4*distance))

	from := math.Floor((float64(time) - d) / 2)
	to := math.Ceil((float64(time) + d) / 2)

	return int(to) - int(from) - 1
}
