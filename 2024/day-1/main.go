package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1

// input.txt

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("Error while reading input file. Check if file exists")
	}

	input := string(data)

	pairs := strings.Split(input, "\n")
	first := []int{}
	second := []int{}

	for _, pair := range pairs {
		locations := strings.Split(pair, "   ")
		firstNum, _ := strconv.Atoi(locations[0])
		secondNum, _ := strconv.Atoi(locations[1])
		first = append(first, firstNum)
		second = append(second, secondNum)
	}

	var part int
	flag.IntVar(&part, "part", 0, "--part=1")
	flag.Parse()

	switch part {
	case 1:
		fmt.Printf("Part 1 result: %d\n", part1(first, second))
	case 2:
		fmt.Printf("Part 2 results: %d\n", part2(first, second))
	default:
		fmt.Printf("%d\n", part1(first, second))
		fmt.Printf("%d\n", part2(first, second))
	}
}

func part1(first, second []int) int {
	sort.Ints(first)
	sort.Ints(second)

	sum := 0
	for i := 0; i < len(first); i++ {
		sum += abs(first[i] - second[i])
	}
	return sum
}

func part2(first, second []int) int {
	occurencesMap := map[int]int{}
	for _, value := range second {
		occurencesMap[value] += 1
	}
	similarityScore := 0
	for _, value := range first {
		similarityScore += value * occurencesMap[value]
	}
	return similarityScore
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
