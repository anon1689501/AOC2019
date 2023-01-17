package main

import (
	"fmt"

	"github.com/anon1689501/AOC2019/lib"
)

func Day2() {
	instructions := lib.ReadInstructions("day2")

	part1input := make([]int, len(instructions))

	copy(part1input, instructions)

	noun := 12
	verb := 2

	part1input[1] = noun
	part1input[2] = verb
	fmt.Println("test2")
	part1Solution, _ := Intcode(part1input)

	fmt.Println("Day 2 Part 1:", part1Solution)

	part2input := make([]int, len(instructions))
	copy(part2input, instructions)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copy(part2input, instructions)
			part2input[1] = i
			part2input[2] = j
			part2Solution, _ := Intcode(part2input)
			if part2Solution == 19690720 {
				fmt.Println("Day 2 Part 2:", 100*i+j)
				break
			}
		}
	}

}
