package main

import "fmt"

/* Go has only one looping construct, the `for` loop */
func main() {
	sum := 0

	/* The basic `for` loop has three components separated by semicolons:
		the init statement: executed before the first iteration
		the condition expression: evaluated before every iteration
		the post statement: executed at the end of every iteration

	The init statement will often be a short variable declaration, and the
	variables declared there are visible only in the scope of the `for`
	statement.

	The loop will stop iterating once the boolean condition evaluates to `false`	   
	*/
	for i := 0; i < 10; i++ { // Braces are always required
		sum += i
	}
	fmt.Println(sum)
}
