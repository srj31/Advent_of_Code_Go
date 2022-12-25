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
		l, r := input[:len(input)/2], input[len(input)/2:]
		for _, c := range l {
			if strings.Contains(r, string(c)) {
				ans += CalculatePriority(c)
				break
			}
		}
	}
	return ans
}

func part_2(inputs []string) int {
	ans := 0
	for i := 0; i < len(inputs); i += 3 {
		l, m, r := inputs[i], inputs[i+1], inputs[i+2]
		for _, c := range l {
			if strings.Contains(m, string(c)) && strings.Contains(r, string(c)) {
				ans += CalculatePriority(c)
				break
			}
		}
	}
	return ans
}

func CalculatePriority(ch rune) int {
	if ch >= 'A' && ch <= 'Z' {
		return int(ch-'A') + 27
	}
	return int(ch-'a') + 1
}
