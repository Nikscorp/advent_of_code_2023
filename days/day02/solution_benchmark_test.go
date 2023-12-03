package main

import "testing"

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

// BenchmarkPart1-10    	    2589	    464993 ns/op	  116329 B/op	    5423 allocs/op
// BenchmarkPart1-10    	    2826	    423393 ns/op	  108399 B/op	    5032 allocs/op
// BenchmarkPart1-10    	   14640	     81025 ns/op	   28312 B/op	     556 allocs/op
// BenchmarkPart1-10    	   15897	     73226 ns/op	   23112 B/op	     447 allocs/op

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}

// BenchmarkPart2-10    	   12768	     93107 ns/op	   27416 B/op	     546 allocs/op
