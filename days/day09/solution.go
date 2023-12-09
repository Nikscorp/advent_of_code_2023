package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0

	for _, line := range strings.Split(input, "\n") {
		numbersS := strings.Split(line, " ")
		numbers := make([]int, len(numbersS))
		for i := range numbersS {
			numbers[i] = toInt(numbersS[i])
		}

		res += findNextRec(numbers)
	}

	return res
}

func part2(input string) any {
	res := 0

	for _, line := range strings.Split(input, "\n") {
		numbersS := strings.Split(line, " ")
		numbers := make([]int, len(numbersS))
		for i := range numbersS {
			numbers[i] = toInt(numbersS[i])
		}

		res += findPrevRec(numbers)
	}

	return res
}

func findNextRec(numbers []int) int {
	diffs := make([]int, len(numbers)-1)
	hasNonZero := false

	for i := 1; i < len(numbers); i++ {
		diffs[i-1] = numbers[i] - numbers[i-1]
		if diffs[i-1] != 0 {
			hasNonZero = true
		}
	}

	if !hasNonZero {
		return numbers[len(numbers)-1]
	}

	return numbers[len(numbers)-1] + findNextRec(diffs)
}

func findPrevRec(numbers []int) int {
	diffs := make([]int, len(numbers)-1)
	hasNonZero := false

	for i := 1; i < len(numbers); i++ {
		diffs[i-1] = numbers[i] - numbers[i-1]
		if diffs[i-1] != 0 {
			hasNonZero = true
		}
	}

	if !hasNonZero {
		return numbers[0]
	}

	return numbers[0] - findPrevRec(diffs)
}

func toInt(s string) int {
	res, _ := strconv.Atoi(s)

	return res
}
