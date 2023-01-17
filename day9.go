package main

import (
	"fmt"

	"github.com/anon1689501/AOC2019/lib"
)

func Day9() {
	inputNumbers := lib.ReadInstructions("Day9")

	day9Part1Solution, _ := Intcode(inputNumbers, 1)

	fmt.Println("answer", day9Part1Solution)
}

//Part 1
//2030 too low
