package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day7() {
	inputText, err := os.Open("input/day7.txt")
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

	inputNumbersBackup := make([]int, len(inputNumbers))
	copy(inputNumbersBackup, inputNumbers)
	phase := []int{0, 1, 2, 3, 4}
	ampOutput := 0
	maxOutput := 0
	//maxPhase := make([]int, 5)

	for a := range phase {
		for b := range phase {
			for c := range phase {
				for d := range phase {
					for e := range phase {
						if a == b || a == c || a == d || a == e || b == c || b == d || b == e || c == d || c == e || d == e {
							continue
						} else {
							ampOutput = 0
							localPhase := []int{a, b, c, d, e}
							index := 0
							for { //_, v := range localPhase {
								copy(inputNumbers, inputNumbersBackup)
								ampOutput = Intcode(inputNumbers, localPhase[index%5], ampOutput)
								index++
							}
							if ampOutput > maxOutput {
								maxOutput = ampOutput
								//copy(maxPhase, localPhase)

							}
						}
					}
				}

			}
		}
	}
	//fmt.Println(maxPhase)
	fmt.Println(maxOutput)

}

//part 1 9210062719823267053 too high
