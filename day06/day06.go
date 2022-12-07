package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("########## Day 6 ##########")
	fmt.Println("Part 01 answer: ", Part01(input))
	fmt.Println("Part 02 answer: ", Part02(input))
}

func Part01(input string) int {
	return getPacketStart(input, 4)
}

func Part02(input string) int {
	return getPacketStart(input, 14)
}

func getPacketStart(input string, amount int) int {
	for i := 0; i < len(input)-amount; i++ {
		if isUnique(input[i : i+amount]) {
			return i + amount
		}
	}
	return 0
}

func isUnique(line string) bool {
	m := make(map[rune]bool)
	for _, c := range line {
		m[c] = false
	}
	return len(line) == len(m)
}
