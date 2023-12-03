package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running part %d...\n", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println(ans)
	} else {
		ans := part2(input)
		fmt.Println(ans)
	}
}
