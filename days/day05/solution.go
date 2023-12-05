package main

import (
	"cmp"
	"math"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type mapInterval struct {
	fromStart, fromEnd int
	toStart, toEnd     int
}

func part1(input string) any {
	groups := strings.Split(input, "\n\n")

	seeds := make([]int, 0)
	for _, seed := range strings.Fields(groups[0][len("seeds: "):]) {
		seeds = append(seeds, toInt(seed))
	}

	mappingPipeline := make([][]mapInterval, 0, len(groups)-1)
	for _, group := range groups[1:] {
		mappingPipeline = append(mappingPipeline, mappingFromS(group))
	}

	res := math.MaxInt
	for _, seed := range seeds {
		resSeed := seed
		for _, mapping := range mappingPipeline {
			resSeed = doMapping(resSeed, mapping)
		}
		res = min(res, resSeed)
	}

	return res
}

func toInt(s string) int {
	res, _ := strconv.Atoi(s)

	return res
}

func mappingFromS(s string) []mapInterval {
	res := make([]mapInterval, 0)

	for _, s := range strings.Split(s, "\n")[1:] {
		fields := strings.Fields(s)
		dst := toInt(fields[0])
		src := toInt(fields[1])
		size := toInt(fields[2])

		res = append(res, mapInterval{
			fromStart: src,
			fromEnd:   src + size - 1,
			toStart:   dst,
			toEnd:     dst + size - 1,
		})
	}
	slices.SortFunc(res, func(a, b mapInterval) int {
		return cmp.Compare(a.fromStart, b.fromStart)
	})

	return res
}

func doMapping(src int, mapping []mapInterval) int {
	pos, exact := slices.BinarySearchFunc(mapping, mapInterval{
		fromStart: src,
	}, func(mi1, mi2 mapInterval) int {
		return cmp.Compare(mi1.fromStart, mi2.fromStart)
	})

	if !exact {
		pos -= 1
	}

	if pos == -1 {
		return src
	}

	interval := mapping[pos]
	if src < interval.fromStart || src > interval.fromEnd {
		return src
	}

	delta := src - interval.fromStart

	return interval.toStart + delta
}

// TODO: now it runs for about ~8m in my mbp.
// should find better approach to do that
func part2(input string) any {
	debug.SetGCPercent(-1)

	groups := strings.Split(input, "\n\n")

	mappingPipeline := make([][]mapInterval, 0, len(groups)-1)
	for _, group := range groups[1:] {
		mappingPipeline = append(mappingPipeline, mappingFromS(group))
	}

	res := math.MaxInt

	seedsS := strings.Fields(groups[0][len("seeds: "):])

	seedCh := make(chan int, 10000)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < len(seedsS); i += 2 {
			seed := toInt(seedsS[i])
			n := toInt(seedsS[i+1])
			for j := 0; j < n; j++ {
				seed++
				seedCh <- seed
			}
		}

		close(seedCh)
	}()

	wgRes := sync.WaitGroup{}
	resCh := make(chan int, 10000)
	for i := 0; i < 8; i++ {
		wgRes.Add(1)

		go func() {
			defer wgRes.Done()
			res := math.MaxInt
			for seed := range seedCh {
				resSeed := seed
				for _, mapping := range mappingPipeline {
					resSeed = doMapping(resSeed, mapping)
				}

				res = min(res, resSeed)
			}

			resCh <- res
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for cur := range resCh {
			res = min(res, cur)
		}
	}()

	wgRes.Wait()
	close(resCh)
	wg.Wait()

	return res
}
