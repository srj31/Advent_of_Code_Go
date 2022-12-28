package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseInput(input string) []string {
	return strings.Split(input, "\n")
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := ParseInput(string(input))

	fmt.Println("Part 1: ", part1(inputs))
	fmt.Println("Part 2: ", part2(inputs))
}

func convertToValues(line string) [4]int {
	re := regexp.MustCompile(`[-,]`)
	pairs := re.Split(line, -1)
	var values [4]int
	for i, pair := range pairs {
		values[i], _ = strconv.Atoi(pair)
	}
	return values
}

func part1(inputs []string) int {
	ans := 0
	for _, line := range inputs {
		values := convertToValues(line)
		if isFullyContainedInside([]int{values[0], values[1]}, []int{values[2], values[3]}) {
			ans++
		}
	}
	return ans
}

func part2(inputs []string) int {
	ans := 0
	for _, line := range inputs {
		values := convertToValues(line)
		if isOverlapping([]int{values[0], values[1]}, []int{values[2], values[3]}) {
			ans++
		}
	}
	return ans
}

func isFullyContainedInside(pair1 []int, pair2 []int) bool {
	return (pair1[0] <= pair2[0] && pair1[1] >= pair2[1]) || (pair1[0] >= pair2[0] && pair1[1] <= pair2[1])
}

func isOverlapping(pair1 []int, pair2 []int) bool {
	return !(pair1[0] > pair2[1] || pair1[1] < pair2[0])
}
