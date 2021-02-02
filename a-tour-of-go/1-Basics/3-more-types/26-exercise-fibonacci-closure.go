package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonaci() func() int {
	fibNumber := 1
	return func(x int) int {
		fibNumber
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
