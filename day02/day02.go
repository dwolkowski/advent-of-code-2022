package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string
var scoresPartOne = map[string]int{"A X": 1 + 3, "B X": 1 + 0, "C X": 1 + 6, "A Y": 2 + 6, "B Y": 2 + 3, "C Y": 2 + 0, "A Z": 3 + 0, "B Z": 3 + 6, "C Z": 3 + 3}
var scoresPartTwo = map[string]int{"A X": 3 + 0, "B X": 1 + 0, "C X": 2 + 0, "A Y": 1 + 3, "B Y": 2 + 3, "C Y": 3 + 3, "A Z": 2 + 6, "B Z": 3 + 6, "C Z": 1 + 6}

func main() {
	fmt.Println("########## Day 2 ##########")
	fmt.Println("Part 01 answer: ", Part01(input))
	fmt.Println("Part 02 answer: ", Part02(input))
}

func Part01(input string) int {
	return getScore(input, scoresPartOne)
}

func Part02(input string) int {
	return getScore(input, scoresPartTwo)
}

func getScore(lines string, scores map[string]int) int {
	score := 0
	for _, line := range strings.Split(lines, "\n") {
		score += scores[strings.TrimSpace(line)]
	}
	return score
}
