package main

import "fmt"

// question: given a number, determine that number's factorial

func main() {
	fmt.Println(recursiveFactorial(5))
	fmt.Println(proceduralFactorial(5))
}

func recursiveFactorial(input int) int{
	if input == 0 {
		return 1
	}

	return input * recursiveFactorial(input - 1)
}

func proceduralFactorial(input int) int {
	if input == 0 {
		return 1
	}

	var accumulator int
	for i := input; i != 0; i-- {
		if i == 0 {
			continue
		}
		if i == input {
			accumulator = i
			continue
		}

		accumulator = accumulator * i
	}

	return accumulator
}