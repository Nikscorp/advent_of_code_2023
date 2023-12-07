package main

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type HandType uint8

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type CardType uint8

const (
	Joker CardType = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	J
	Q
	K
	A
)

type Hand []CardType

type HandWithType struct {
	rawHand  string
	hand     Hand
	handType HandType
	bid      int
}

func (h *HandWithType) Compare(other *HandWithType) int {
	if h.handType == other.handType {
		for i, selfCard := range h.hand {
			if selfCard != other.hand[i] {
				return cmp.Compare(selfCard, other.hand[i])
			}
		}

		return 0
	}

	return cmp.Compare(h.handType, other.handType)
}

func part1(input string) any {
	splitted := strings.Split(input, "\n")
	handsWithType := make([]*HandWithType, 0, len(splitted))

	for _, line := range splitted {
		rawHand, bid, _ := strings.Cut(line, " ")
		handsWithType = append(handsWithType, &HandWithType{
			rawHand:  rawHand,
			hand:     hand(rawHand, false),
			handType: handType(rawHand, false),
			bid:      toInt(bid),
		})
	}

	slices.SortFunc(handsWithType, func(a, b *HandWithType) int {
		return a.Compare(b)
	})

	res := 0

	for i, hand := range handsWithType {
		res += hand.bid * (i + 1)
	}

	return res
}

func part2(input string) any {
	splitted := strings.Split(input, "\n")
	handsWithType := make([]*HandWithType, 0, len(splitted))

	for _, line := range splitted {
		rawHand, bid, _ := strings.Cut(line, " ")
		handsWithType = append(handsWithType, &HandWithType{
			rawHand:  rawHand,
			hand:     hand(rawHand, true),
			handType: handType(rawHand, true),
			bid:      toInt(bid),
		})
	}

	slices.SortFunc(handsWithType, func(a, b *HandWithType) int {
		return a.Compare(b)
	})

	res := 0

	for i, hand := range handsWithType {
		res += hand.bid * (i + 1)
	}

	return res
}

func hand(s string, considerJokers bool) Hand {
	var res = make(Hand, 5)

	for i := 0; i < len(s); i++ {
		v := s[i]
		switch {
		case v >= '2' && v <= '9':
			res[i] = CardType(v - '2' + uint8(Two))
		case v == 'T':
			res[i] = Ten
		case v == 'J':
			if considerJokers {
				res[i] = Joker
			} else {
				res[i] = J
			}
		case v == 'Q':
			res[i] = Q
		case v == 'K':
			res[i] = K
		case v == 'A':
			res[i] = A
		default:
			panic("unknown card")
		}
	}

	return res
}

func handType(s string, considerJokers bool) HandType {
	m := make(map[byte]int)
	jokerCnt := 0

	for i := 0; i < len(s); i++ {
		if considerJokers && s[i] == 'J' {
			jokerCnt++
			continue
		}
		m[s[i]]++
	}

	if considerJokers && jokerCnt > 0 {
		var keyOfMaxV byte
		maxV := 0
		for k, v := range m {
			if v > maxV {
				maxV = v
				keyOfMaxV = k
			}
		}

		m[keyOfMaxV] += jokerCnt
	}

	switch len(m) {
	case 1:
		return FiveOfAKind
	case 5:
		return HighCard
	case 4:
		return OnePair
	}

	maxV := 0
	for _, v := range m {
		maxV = max(maxV, v)
	}

	switch len(m) {
	case 2:
		if maxV == 4 {
			return FourOfAKind
		}
		return FullHouse
	default:
		if maxV == 3 {
			return ThreeOfAKind
		}
		return TwoPair
	}
}

func toInt(s string) int {
	res, _ := strconv.Atoi(s)

	return res
}
