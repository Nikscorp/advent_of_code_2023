package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		line, numsS, _ := strings.Cut(line, " ")
		numsSplitted := strings.Split(numsS, ",")
		nums := make([]int, len(numsSplitted))
		for i, ns := range numsSplitted {
			nums[i], _ = strconv.Atoi(ns)
		}

		line = strings.Replace(line, "..", ".", -1)
		res += dp(nums, line, 0, 0, false, make(map[memoKey]int))
	}

	return res
}

type memoKey struct {
	lineIndex int
	numsIndex int
	prevHash  bool
}

func dp(nums []int, line string, lineIndex int, numsIndex int, prevHash bool, memo map[memoKey]int) (res int) {
	if lineIndex >= len(line) {
		if numsIndex >= len(nums) {
			return 1
		}

		return 0
	}

	if numsIndex >= len(nums) {
		if strings.Contains(line[lineIndex:], "#") {
			return 0
		}

		return 1
	}

	memoKey := memoKey{lineIndex, numsIndex, prevHash}
	if v, ok := memo[memoKey]; ok {
		return v
	}

	defer func() {
		memo[memoKey] = res
	}()

	if line[lineIndex] == '.' {
		return dp(nums, line, lineIndex+1, numsIndex, false, memo)
	}

	if line[lineIndex] == '#' {
		if prevHash {
			return 0
		}

		if lineIndex+nums[numsIndex] > len(line) {
			return 0
		}

		if strings.Contains(line[lineIndex:lineIndex+nums[numsIndex]], ".") {
			return 0
		}

		return dp(nums, line, lineIndex+nums[numsIndex], numsIndex+1, true, memo)
	}

	// line[lineIndex] == '?'

	// assume it is a '.'
	res = dp(nums, line, lineIndex+1, numsIndex, false, memo)

	// assume it is a '#'
	if lineIndex+nums[numsIndex] <= len(line) && !strings.Contains(line[lineIndex:lineIndex+nums[numsIndex]], ".") && !prevHash {
		res += dp(nums, line, lineIndex+nums[numsIndex], numsIndex+1, true, memo)
	}

	return res
}

func part2(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		line, numsS, _ := strings.Cut(line, " ")
		numsSplitted := strings.Split(numsS, ",")
		nums := make([]int, len(numsSplitted))
		for i, ns := range numsSplitted {
			nums[i], _ = strconv.Atoi(ns)
		}

		line = strings.Join(repeat([]string{line}, 5), "?")
		line = strings.Replace(line, "..", ".", -1)
		nums = repeat(nums, 5)

		res += dp(nums, line, 0, 0, false, make(map[memoKey]int))
	}

	return res
}

func repeat[T any](slice []T, n int) []T {
	result := make([]T, 0, len(slice)*n)
	for i := 0; i < n; i++ {
		result = append(result, slice...)
	}

	return result
}
