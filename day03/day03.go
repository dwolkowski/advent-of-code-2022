package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("########## Day 3 ##########")
	fmt.Println("Part 01 answer: ", Part01(input))
	fmt.Println("Part 02 answer: ", Part02(input))
}

func Part01(input string) int {
	sum := 0
	for _, line := range strings.Fields(input) {
		sum += getPriorityForLine(line)
	}
	return sum
}

func Part02(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i += 3 {
		for _, c := range lines[i] {
			if strings.Contains(lines[i+1], string(c)) && strings.Contains(lines[i+2], string(c)) {
				sum += getPriority(c)
				break
			}
		}
	}
	return sum
}

func getPriorityForLine(line string) int {
	left, right := line[:len(line)/2], line[len(line)/2:]
	for _, c := range left {
		if strings.Contains(right, string(c)) {
			return getPriority(c)
		}
	}
	return 0
}

func getPriority(char int32) int {
	if char >= 65 && char <= 90 {
		return int(char%32 + 26)
	} else {
		return int(char % 32)
	}
}
