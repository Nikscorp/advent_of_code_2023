package main

import (
	"strings"
)

func part1(input string) any {
	res := 0

	for _, line := range strings.Split(input, "\n") {
		digit := 0

		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				digit += int(line[i] - '0')
				break
			}
		}
		digit *= 10

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				digit += int(line[i] - '0')
				break
			}
		}

		res += digit
	}

	return res
}

func part2(input string) any {
	res := 0

	var m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range strings.Split(input, "\n") {
		digit := 0

	L1:
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				digit += int(line[i] - '0')
				break
			}

			for k, v := range m {
				if strings.HasPrefix(line[i:], k) {
					digit += v
					break L1
				}
			}
		}

		digit *= 10

	L2:
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				digit += int(line[i] - '0')
				break
			}

			for k, v := range m {
				if strings.HasSuffix(line[:i+1], k) {
					digit += v
					break L2
				}
			}
		}

		res += digit
	}

	return res
}
