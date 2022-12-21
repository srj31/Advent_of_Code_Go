package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

	fmt.Println("Part 1: ", part_1(inputs))
	fmt.Println("Part 2: ", part_2(inputs))
}

func part_1(inputs []string) int {
	ans := 0
	cur := 0
	for _, cal := range inputs {
		if cal != "" {
			calInInt, _ := strconv.Atoi(cal)
			cur += calInInt
			continue
		}
		if cur > ans {
			ans = cur
		}
		cur = 0
	}

	return ans
}

func part_2(inputs []string) int {
	ans := 0
	cur := 0

	var calPerElf []int
	for _, cal := range inputs {
		if cal != "" {
			calInInt, _ := strconv.Atoi(cal)
			cur += calInInt
			continue
		}
		calPerElf = append(calPerElf, cur)
		cur = 0
	}

	sort.Slice(calPerElf, func(i, j int) bool {
		return calPerElf[i] > calPerElf[j]
	})

	for _, cal := range calPerElf[0:3] {
		ans += cal
	}
	return ans
}
