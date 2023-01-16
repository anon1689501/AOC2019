package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	inputText, err := os.Open("input/day5.txt")
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

	Intcode(inputNumbers)

}
