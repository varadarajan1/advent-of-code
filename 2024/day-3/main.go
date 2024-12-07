package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("Error while reading input file. Check if file exists")
	}

	input := string(data)

	var part int
	flag.IntVar(&part, "part", 0, "--part=1")
	flag.Parse()

	switch part {
	case 1:
		fmt.Printf("Part 1 result: %d\n", part1(input))
	case 2:
		fmt.Printf("Part 2 results: %d\n", part2(input))
	default:
		fmt.Printf("%d\n", part1(input))
		fmt.Printf("%d\n", part2(input))
	}
}

func part1(input string) int {
	regex, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	if err != nil {
		panic(err)
	}
	matches := regex.FindAllStringSubmatch(input, -1)
	res := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		res += num1 * num2
	}
	return res
}

func part2(input string) int {
	regex, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)|don't\\(\\)|do\\(\\)")
	if err != nil {
		panic(err)
	}
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0
	skip := false
	for _, match := range matches {
		if strings.HasPrefix(match[0], "don't()") {
			skip = true
			continue
		}

		if strings.HasPrefix(match[0], "do()") {
			skip = false
			continue
		}
		if !skip {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			sum += x * y
		}
	}
	return sum
}
