package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	previousNumber := 0
	currentNumber := 1

	return func() int {
		oldPreviousNumber := previousNumber
		newPreviousNumber := currentNumber
		currentNumber += previousNumber
		previousNumber = newPreviousNumber
		return oldPreviousNumber
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
