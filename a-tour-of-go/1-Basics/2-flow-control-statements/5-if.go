package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 { // No parenthesis surrounding expression. Braces required
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
