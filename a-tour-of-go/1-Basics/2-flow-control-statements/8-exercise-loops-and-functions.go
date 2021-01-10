package main

import (
   "fmt"
   "math"
)

/* As a way to play with functions and loops, let's implement a square root
   function: given a number `x`, we want to find the number `z` for which z² is
   most nearly `x`.

   Computers typically compute the square root of `x` using a loop. Starting with
   some guess `z`, we can adjust `z` based on how close `z²` is to `x`,
   producing a better guess:

z -= (z*z - x) / (2*z)

   Repeating this adjustment makes the guess better and better until we reach an
   answer that is as close to the actual square root as can be.

   Implement this in the `func sqrt` provided. A decent starting guess for `z` is
   1, no matter what the input. To begin with, repeat the calculation 10 times
   and print each `z` along the way. See how close you get to the answer for
   various values of `x` (1, 2, 3...) and how quickly the guess improves.

   Next, change the loop condition to stop once the value has stopped changing
   (or only changes by a very small amount). See if that's more or fewer than
   10 iterations. Try other initial guesses for `z`, like `x`, or `x/2`. How
   close are your function's results to the math.Sqrt in the standard library? */
func Sqrt(x float64, z float64) float64 {
   const delta = 0.000000000000001 // A very small number
	var previous_guess float64

   iteration := 0 
	for {
      iteration += 1

		z -= (z*z - x) / (2 * z) // Newton's formula

      /* If `z` didn't change, or the difference between previous `z` and the new
      `z` is too small, this is our guess */
      if previous_guess == z || math.Abs(previous_guess-z) < delta {
         fmt.Printf("took %d iterations.\n", iteration)
			return z
		}

      previous_guess = z
   }
}

func main() {
   // What numbers should we find the square root for
   for x := 1; x <= 20; x++ {
      fmt.Println("------------------------------------")
      fmt.Printf("Calculating Sqrt(%d):\n", x)

      // Initial guesses: 1, x, x/2
      initial_guesses := []float64{1.0, float64(x), float64(float64(x)/float64(2))}
      
      for _, initial_guess := range initial_guesses {
         fmt.Printf("Initial guess: %f\n", initial_guess)
         
         // Calculate root by Newton's method
         guess := Sqrt(float64(x), initial_guess)
         // Calculate root with standard library
         actual_root := math.Sqrt(float64(x))
         
         fmt.Printf("Root guess : %g | ", guess)
         fmt.Printf("Actual root: %g\n", actual_root)
         // Print how different are Newton's and Standard library root
         fmt.Printf("Difference: %g\n\n", math.Abs(guess - actual_root))
      }
   }
}
