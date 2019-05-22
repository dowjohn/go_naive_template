package main

import "fmt"

// question: given an index, determine the fibonacci value which falls on that index. 0,1,2,3,4,5 indexes correspond with 0,1,1,2,3,5 fibonacci
//0,1,1,2
func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))

	fmt.Println(fibAgain(0))
	fmt.Println(fibAgain(1))
	fmt.Println(fibAgain(2))
	fmt.Println(fibAgain(3))
	fmt.Println(fibAgain(4))
	fmt.Println(fibAgain(5))
	fmt.Println(fibAgain(6))
	fmt.Println(fibAgain(7))
	fmt.Println(fibAgain(8))
}

// recursive solution - dont use a big number :/
func fib(index int) int {
	if index == 0 || index == 1 {
		return index
	}

	return fib(index - 2) + fib(index - 1)
}

// procedural solution
func fibAgain(index int) int {
	zero := 0
	one := 1
	two := 2
	if index == zero {
		return zero
	}

	if index == one {
		return one
	}

	if index == two {
		return one
	}

	var sum int
	previousLastValue := 1
	lastValue := 1
	for i := 2; i < index; i++ {
		sum = previousLastValue + lastValue
		previousLastValue = lastValue
		lastValue = sum
	}

	return sum
}