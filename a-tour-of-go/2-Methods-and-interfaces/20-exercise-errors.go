package main

import (
	"fmt"
	"math"
)

/* Copy your `Sqrt` function from the earlier exercise and modify it to return an
   `error` value.

   `Sqrt` should return a non-nil error value when given a negative number, as it
   doesn't support complex numbers.

   Create a new type
type ErrNegativeSqrt float64

   And make it an `error` by giving it a
func (e ErrNegativeSqrt) Error() string

   method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative
   number: -2".

   Note: A call to `fmt.Sprint(e)` inside the `Error` method will send the program
   into an infinite loop. You can avoid this by converting `e` first:
   `fmt.Sprint(float64(e))`. Why?

fmt.Sprint(e) will call e.Error() to convert the value e to a string.
If the Error() method calls fmt.Sprint(e), then the program recurses until out
of memory.
You can break the recursion by converting the e to a value without a 
String or Error method.

   Change your `Sqrt` function to return an `ErrNegativeSqrt` value when given
   a negative number.
*/

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := float64(x/2) // Initial guess
	const delta = 0.000000000000001 // A very small number
	var previous_guess float64

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	iteration := 0
	for {
		iteration += 1

		z -= (z*z - x) / (2 * z) // Newton's formula

		/* If `z` didn't change, or the difference between previous `z` and the new
		   `z` is too small, this is our guess */
		if previous_guess == z || math.Abs(previous_guess-z) < delta {
			fmt.Printf("took %d iterations.\n", iteration)
			return z, nil
		}

		previous_guess = z
	}
}

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-2))
}
