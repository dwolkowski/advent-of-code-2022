package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("########## Day 1 ##########")
	fmt.Println("Part 01 answer: ", Part01(input))
	fmt.Println("Part 02 answer: ", Part02(input))
}

func Part01(input string) int {
	return getCaloriesSlice(input)[0]
}

func Part02(input string) int {
	kcal := getCaloriesSlice(input)
	return kcal[0] + kcal[1] + kcal[2]
}

func getCaloriesSlice(input string) []int {
	actual := 0
	var calories []int

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			calories = append(calories, actual)
			actual = 0
		} else {
			elfCalories, _ := strconv.Atoi(line)
			actual += elfCalories
		}
	}
	calories = append(calories, actual)

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	return calories
}
