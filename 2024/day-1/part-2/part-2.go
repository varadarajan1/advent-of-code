package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1#part2
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

	occurencesMap := map[int]int{}
	for _, value := range second {
		occurencesMap[value] += 1
	}
	similarityScore := 0
	for _, value := range first {
		similarityScore += value * occurencesMap[value]
	}
	fmt.Printf("Similarity Score:%d", similarityScore)
}
