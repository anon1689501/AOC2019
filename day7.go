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
	phase := []int{5, 6, 7, 8, 9}
	ampOutput := 0
	maxOutput := 0
	maxPhase := make([]int, 5)

	for _, a := range phase {
		for _, b := range phase {
			for _, c := range phase {
				for _, d := range phase {
					for _, e := range phase {
						if a == b || a == c || a == d || a == e || b == c || b == d || b == e || c == d || c == e || d == e {
							continue
						} else {
							ampOutput = 0
							complete := false
							localPhase := []int{a, b, c, d, e}
							index := 0
							phaseInput := 0
							//fmt.Println(localPhase)
							for { //_, v := range localPhase {
								copy(inputNumbers, inputNumbersBackup)
								if index < 5 {
									phaseInput = localPhase[index]
								} else {
									phaseInput = ampOutput
								}
								ampOutput, complete = Intcode(inputNumbers, phaseInput, ampOutput)
								index++
								if complete {
									break
								}
							}
							if ampOutput > maxOutput {
								fmt.Println(index)
								maxOutput = ampOutput
								copy(maxPhase, localPhase)

							}
						}
					}
				}

			}
		}
	}
	fmt.Println(maxPhase)
	fmt.Println(maxOutput)

}

//part 1 9210062719823267053 too high
