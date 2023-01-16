package main

import "fmt"

func main() {
	Day7()

}
func Intcode(input []int, params ...int) (int, bool) {

	//for day 1 uncomment below
	// if len(params) == 2 {
	// 	input[1] = params[0] //noun
	// 	input[2] = params[1] //verb
	// }

	paramCurrent := 0

	solution := 0
	finished := false
	//out:
	for i := 0; i < len(input); {
		//fmt.Println("i", i)
		instruction := input[i] % 100
		firstParamMode := (input[i] / 100) % 10   // 0 = position mode (day2 style)
		secondParamMode := (input[i] / 1000) % 10 // 1 = immediate
		//thirdParamMode := input[i] / 10000

		arg1 := 0
		arg2 := 0
		//arg3 := 0
		//fmt.Println(input)

		//fmt.Println(i, input[i], input[i+1], input[i+2])

		switch instruction { //if first arg is needed
		case 1, 2, 4, 5, 6, 7, 8:
			//fmt.Println(input[i])
			if firstParamMode == 1 {
				arg1 = input[i+1] //immediate mode
			} else {
				arg1 = input[input[i+1]]
			}
		}

		switch instruction { //if second arg is needed
		case 1, 2, 5, 6, 7, 8:
			if secondParamMode == 1 {
				arg2 = input[i+2] //immediate mode
			} else {
				arg2 = input[input[i+2]]
			}
		}

		switch instruction {

		case 1: //add two numbers
			//fmt.Println("index: ", i, arg1, "+", arg2, "at", input[i+3])
			input[input[i+3]] = arg1 + arg2
			i += 4

		case 2: //multiply two numbers
			//fmt.Println("index: ", i, arg1, "*", arg2, "at", input[i+3])
			input[input[i+3]] = arg1 * arg2
			i += 4

		case 3: //read a value from params
			//fmt.Println("print register", i, input[i+1])
			input[input[i+1]] = params[paramCurrent]
			paramCurrent++
			i += 2

		case 4: //print a value
			//fmt.Println(arg1)
			solution = solution*10 + arg1
			i += 2

		case 5: //jump if not zero
			if arg1 != 0 {
				i = arg2
				//fmt.Println(i)
			} else {
				i += 3
			}

		case 6: //jump if zero
			if arg1 == 0 {
				i = arg2
				//fmt.Println(i)
			} else {
				i += 3
			}

		case 7: //less than
			if arg1 < arg2 {
				input[input[i+3]] = 1
			} else {
				input[input[i+3]] = 0
			}
			i += 4

		case 8: //equals
			if arg1 == arg2 {
				input[input[i+3]] = 1
			} else {
				input[input[i+3]] = 0
			}
			i += 4

		case 99: //exit
			finished = true
			//i = len(input)
			//break out
			// default:
			// 	fmt.Println("unknown command ", input[i])
		}

	}
	fmt.Println("")
	return solution, finished
}
