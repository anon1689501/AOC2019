package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2() {

	inputText, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	inputNumbers := make([]int, 0)

	for scanner.Scan() {
		stringSlice := strings.Split(scanner.Text(), ",")
		for _, v := range stringSlice {
			number, _ := strconv.Atoi(v)
			inputNumbers = append(inputNumbers, number)
		}
	}

	part1input := make([]int, len(inputNumbers))
	copy(part1input, inputNumbers)

	fmt.Println("Day 2 Part 1:", Intcode(part1input, 12, 2))

	part2input := make([]int, len(inputNumbers))
	copy(part2input, inputNumbers)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copy(part2input, inputNumbers)
			if Intcode(part2input, i, j) == 19690720 {
				fmt.Println("Day 2 Part 2:", 100*i+j)
				break
			}
		}
	}

}
