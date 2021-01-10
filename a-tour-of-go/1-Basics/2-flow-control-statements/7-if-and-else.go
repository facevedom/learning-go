package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	/* Variables declred inside an `if` short statement are also available 
	   inside any of the `else` blocks. */
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	// Both calls to `pow` return before the call to `fmt.Println` in `main` begins
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
