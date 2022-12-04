package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("########## Day 4 ##########")
	fmt.Println("Part 01 answer: ", Part01(input))
	fmt.Println("Part 02 answer: ", Part02(input))
}

func Part01(input string) int {
	var left, right section
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		fmt.Sscanf(line, "%d-%d,%d-%d", &left.start, &left.end, &right.start, &right.end)

		if doesFullyContain(left, right) {
			sum++
		}
	}

	return sum
}

func Part02(input string) int {
	var left, right section
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		fmt.Sscanf(line, "%d-%d,%d-%d", &left.start, &left.end, &right.start, &right.end)

		if doesOverlap(left, right) {
			sum++
		}
	}

	return sum
}

func doesFullyContain(left section, right section) bool {
	return (left.start <= right.start && left.end >= right.end) ||
		(right.start <= left.start && right.end >= left.end)
}

func doesOverlap(left section, right section) bool {
	return left.end >= right.start && right.end >= left.start
}

type section struct{ start, end int }
