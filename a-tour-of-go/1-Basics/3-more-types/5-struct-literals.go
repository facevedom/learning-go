package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	/* A struct literal denotes a newly allocated struct value by listing the
	   values of its fields. */
	v1 = Vertex{1, 2}  // has type Vertex
	/* You can list just a subset of fields by using the `Name:` syntax.
	   (And the order of named fields is irrelevant.) */
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	// The special prefix `&` returns a pointer to the struct value
	p = &Vertex{1, 2}  // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
