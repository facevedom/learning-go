package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Struct fields can be accessed through a struct pointer.
func main() {
	v := Vertex{1, 2}
	p := &v
	/* To access te field `X` of a struct when we have the struct pointer `p` we
	   could write `(*p).X`. However, that notation is cumbersome, so the
	   language permits us instead to write just `p.X`, without the explicit
	   dereference */
	p.X = 1e9
	fmt.Println(v)
}
