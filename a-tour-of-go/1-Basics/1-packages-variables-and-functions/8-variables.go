package main

import "fmt"

/* The `var` statement declares a list of variables; as in function argument
   lists, the type is last. */
var c, python, java bool

func main() {
    /* A `var` statement can be at package or function level. We see both in this
       example */
	var i int
	fmt.Println(i, c, python, java)
}
