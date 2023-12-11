package main

import "testing"

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

// 256_683_146
// 119_800_324
// 107_954_162

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
