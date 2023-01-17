package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInstructions(day string) []int {

	inputString := "./input/" + day + ".txt"
	inputText, err := os.Open(inputString)
	if err != nil {
		log.Fatal(err)
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

	return inputNumbers

}
