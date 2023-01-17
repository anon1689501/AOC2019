package main

import (
	"fmt"

	"github.com/anon1689501/AOC2019/lib"
)

func Day5() {
	inputNumbers := lib.ReadInstructions("day5")
	inputNumbersBackup := make([]int, len(inputNumbers))
	copy(inputNumbersBackup, inputNumbers)

	day5Part1Solution, _ := Intcode(inputNumbers, 1)
	fmt.Println(day5Part1Solution)

	day5Part2Solution, _ := Intcode(inputNumbersBackup, 5)
	fmt.Println(day5Part2Solution)

}
