package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("########## Day 5 ##########")
	fmt.Println("Part 01 answer: ", Part01(input))
	fmt.Println("Part 02 answer: ", Part02(input))
}

func Part01(input string) string {
	var amount, from, to int
	for _, line := range strings.Split(input, "\n") {
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
		if err != nil {
			continue
		}
		move(from, to, amount)
	}

	return buildAnswer(stack)
}

func Part02(input string) string {
	var amount, from, to int
	for _, line := range strings.Split(input, "\n") {
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
		if err != nil {
			continue
		}
		moveWithSameOrder(from, to, amount)
	}

	return buildAnswer(stack2)
}

func move(from int, to int, amount int) {
	for i := 0; i < amount; i++ {
		stack[to] = append(stack[to], stack[from][len(stack[from])-1])
		stack[from] = append(stack[from][:len(stack[from])-1])
	}
}

func moveWithSameOrder(from int, to int, amount int) {
	stack2[to] = append(stack2[to], stack2[from][len(stack2[from])-amount:]...)
	stack2[from] = append(stack2[from][:len(stack2[from])-amount])
}

func buildAnswer(anyStack map[int][]string) string {
	s := ""
	for i := 1; i <= len(anyStack); i++ {
		s += anyStack[i][len(anyStack[i])-1]
	}
	return s
}

var stack = map[int][]string{
	1: {"N", "R", "G", "P"},
	2: {"J", "T", "B", "L", "F", "G", "D", "C"},
	3: {"M", "S", "V"},
	4: {"L", "S", "R", "C", "Z", "P"},
	5: {"P", "S", "L", "V", "C", "W", "D", "Q"},
	6: {"C", "T", "N", "W", "D", "M", "S"},
	7: {"H", "D", "G", "W", "P"},
	8: {"Z", "L", "P", "H", "S", "C", "M", "V"},
	9: {"R", "P", "F", "L", "W", "G", "Z"},
}

var stack2 = map[int][]string{
	1: {"N", "R", "G", "P"},
	2: {"J", "T", "B", "L", "F", "G", "D", "C"},
	3: {"M", "S", "V"},
	4: {"L", "S", "R", "C", "Z", "P"},
	5: {"P", "S", "L", "V", "C", "W", "D", "Q"},
	6: {"C", "T", "N", "W", "D", "M", "S"},
	7: {"H", "D", "G", "W", "P"},
	8: {"Z", "L", "P", "H", "S", "C", "M", "V"},
	9: {"R", "P", "F", "L", "W", "G", "Z"},
}
