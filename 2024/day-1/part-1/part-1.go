package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1

// input.txt

func main() {
	data, err := os.ReadFile("./day-1/part-1-input.txt")
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

	sort.Ints(first)
	sort.Ints(second)

	sum := 0
	for i := 0; i < len(pairs); i++ {
		sum += abs(first[i] - second[i])
	}
	fmt.Printf("Total distance: %d", sum)
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
