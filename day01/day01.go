package day01

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func SolvePart1(input <-chan string) int {
	pairs := parse(input)
	lefts := make([]int, len(pairs))
	rights := make([]int, len(pairs))

	for i, pair := range pairs {
		lefts[i] = pair[0]
		rights[i] = pair[1]
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	sum := 0
	for i := range len(lefts) {
		sum += max(lefts[i], rights[i]) - min(lefts[i], rights[i])
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	pairs := parse(input)
	rightCounts := make(map[int]int)

	for _, pair := range pairs {
		rightCounts[pair[1]]++
	}

	sum := 0
	for _, pair := range pairs {
		sum += pair[0] * rightCounts[pair[0]]
	}

	return sum
}

func parse(input <-chan string) [][2]int {
	pairs := make([][2]int, 0)

	for line := range input {
		idx := strings.IndexFunc(line, unicode.IsSpace)
		first := line[:idx]
		second := strings.TrimLeftFunc(line[idx:], unicode.IsSpace)

		left, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(second)
		if err != nil {
			panic(err)
		}

		pairs = append(pairs, [2]int{left, right})
	}

	return pairs
}
