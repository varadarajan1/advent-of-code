package main

import (
	"flag"
	"fmt"
	"os"
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

	rows := strings.Split(input, "\n")

	parsedNums := [][]int{}
	for _, row := range rows {
		nums := strings.Split(row, " ")
		arr := []int{}
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			arr = append(arr, n)
		}
		parsedNums = append(parsedNums, arr)
	}

	var part int
	flag.IntVar(&part, "part", 0, "--part=1")
	flag.Parse()

	switch part {
	case 1:
		fmt.Printf("Part 1 result: %d\n", part1(parsedNums))
	case 2:
		fmt.Printf("Part 2 results: %d\n", part2(parsedNums))
	default:
		fmt.Printf("%d\n", part1(parsedNums))
		fmt.Printf("%d\n", part2(parsedNums))
	}
}

func part1(arr [][]int) int {
	safeLevels := 0
	for _, nums := range arr {
		if isSafe(nums) {
			safeLevels += 1
		}
	}
	return safeLevels
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func part2(arr [][]int) int {
	safeLevels := 0
	for _, nums := range arr {
		if isSafe(nums) {
			safeLevels += 1
			continue
		}

		for i := range nums {
			newArr := []int{}
			for j := 0; j < len(nums); j++ {
				if i != j {
					newArr = append(newArr, nums[j])
				}
			}
			if isSafe(newArr) {
				safeLevels += 1
				break
			}
		}
	}
	return safeLevels
}

func isSafe(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}
	dir := arr[0] - arr[1]
	for idx := 1; idx < len(arr); idx++ {
		diff := arr[idx-1] - arr[idx]
		if abs(diff) < 1 || abs(diff) > 3 || (dir*diff <= 0) {
			return false
		}
	}
	return true
}
