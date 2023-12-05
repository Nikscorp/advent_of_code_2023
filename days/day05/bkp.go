package main

// TODO it's really tough to optimize, should do it later
// func part2(input string) any {
// 	groups := strings.Split(input, "\n\n")

// 	mappingPipeline := make([][]mapInterval, 0, len(groups))
// 	seedsS := strings.Fields(groups[0][len("seeds: "):])
// 	for i := 0; i < len(seedsS); i += 2 {
// 		seedIntFrom, _ := strconv.Atoi(seedsS[i])
// 		n, _ := strconv.Atoi(seedsS[i+1])
// 		mappingPipeline = append(mappingPipeline, []mapInterval{
// 			{
// 				fromStart: seedIntFrom,
// 				fromEnd:   seedIntFrom + n - 1,
// 				toStart:   seedIntFrom,
// 				toEnd:     seedIntFrom + n - 1,
// 				fake:      true,
// 			},
// 		})
// 	}

// 	for _, group := range groups[1:] {
// 		mappingPipeline = append(mappingPipeline, mappingFromS(group))
// 	}

// 	for len(mappingPipeline) > 1 {
// 		for _, interval := range mappingPipeline[0] {
// 			interval.fromStart = interval.toStart
// 			interval.fromEnd = interval.toEnd
// 			interval.fake = true
// 		}
// 		slices.SortFunc(mappingPipeline[0], func(a, b mapInterval) int {
// 			return cmp.Compare(a.fromStart, b.fromStart)
// 		})

// 		intersected := intervalIntersection(mappingPipeline[0], mappingPipeline[1])
// 		intersected = append(intersected, mappingPipeline[0]...)
// 		mappingPipeline = mappingPipeline[1:]
// 		mappingPipeline[0] = intersected
// 	}

// 	res := math.MaxInt
// 	for _, intervals := range mappingPipeline {
// 		if len(intervals) == 0 {
// 			continue
// 		}

// 		for _, interval := range intervals {
// 			if interval.fake {
// 				continue
// 			}
// 			res = min(res, interval.toStart)
// 		}

// 	}

// 	return res
// }

// func intervalIntersection(first, second []mapInterval) []mapInterval {
// 	fp := 0
// 	sp := 0
// 	res := make([]mapInterval, 0)

// 	for fp < len(first) && sp < len(second) {
// 		v, ok := intersects(first[fp], second[sp])
// 		if ok {
// 			res = append(res, v)
// 		}

// 		if first[fp].fromEnd == second[sp].fromEnd {
// 			fp++
// 			sp++
// 		} else if first[fp].fromEnd < second[sp].fromEnd {
// 			fp++
// 		} else {
// 			sp++
// 		}
// 	}

// 	return res
// }

// func intersects(a, b mapInterval) (mapInterval, bool) {
// 	if a.fromStart > b.fromStart {
// 		a, b = b, a
// 	}

// 	if b.fromStart > a.fromEnd {
// 		return mapInterval{}, false
// 	}

// 	res := b
// 	res.fromEnd = min(a.fromEnd, b.fromEnd)
// 	if res.fake {
// 		res.toStart = a.toStart
// 		res.toEnd = a.toEnd
// 		res.fake = false
// 	}

// 	delta := res.fromEnd - res.fromStart
// 	res.toEnd = res.fromEnd + delta

// 	return res, true
// }
