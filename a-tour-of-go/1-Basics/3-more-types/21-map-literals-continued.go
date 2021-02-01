package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

/* If the top-level type is just a type name, you can omit it from the elements
   of the literal. */
var m = map[string]Vertex {
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
