package main

import (
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0
	for _, word := range strings.Split(input, ",") {
		res += hash(word)
	}

	return res
}

type box struct {
	orderedLabels []string
	lenses        map[string]int
}

func newBox() *box {
	return &box{
		orderedLabels: make([]string, 0),
		lenses:        make(map[string]int),
	}
}

func part2(input string) any {
	boxes := make([]*box, 0, 255)
	for i := 0; i < 256; i++ {
		boxes = append(boxes, newBox())
	}

	for _, word := range strings.Split(input, ",") {
		opInd := strings.IndexAny(word, "-=")
		op := word[opInd]
		label := word[:opInd]
		h := hash(label)
		if op == '-' {
			delete(boxes[h].lenses, label)
		} else {
			f, _ := strconv.Atoi(word[opInd+1:])
			if _, ok := boxes[h].lenses[label]; !ok {
				boxes[h].orderedLabels = append(boxes[h].orderedLabels, label)
			}
			boxes[h].lenses[label] = f
		}
	}

	res := 0

	for i, box := range boxes {
		curSlot := 1
		for j, label := range box.orderedLabels {

			v, ok := box.lenses[label]
			if !ok {
				continue
			}
			if slices.Contains(box.orderedLabels[j+1:], label) {
				continue
			}

			res += (i + 1) * curSlot * v
			curSlot++
		}
	}

	return res
}

func hash(word string) int {
	cur := 0
	for i := 0; i < len(word); i++ {
		cur += int(word[i])
		cur *= 17
		cur %= 256
	}

	return cur
}
