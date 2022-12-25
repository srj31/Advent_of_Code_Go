package main

import (
	"fmt"
	"log"
	"os"
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

	fmt.Println("Part 1: ", part_1(inputs))
	fmt.Println("Part 2: ", part_2(inputs))
}

func part_1(inputs []string) int {
	ans := 0
	for _, input := range inputs {
		l := input[0] - 'A'
		r := input[2] - 'X'
		ans += int(r + 1)
		ans += (int((r)-(l)+4) % 3) * 3
	}
	return ans
}

func part_2(inputs []string) int {
	ans := 0
	for _, input := range inputs {
		l := input[0] - 'A'
		r := input[2] - 'X'
		ans += int(r) * 3
		ans += (int(l+r+2) % 3) + 1
	}
	return ans
}
